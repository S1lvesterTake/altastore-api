package login

import (
	"altastore-api/application/helper"
	"altastore-api/application/infrastructure"
	"context"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
	"golang.org/x/crypto/bcrypt"
)

//LoginHandler login handler struct
type LoginHandler struct {
	repository   infrastructure.LoginRepository
	accountRepo  infrastructure.AccountRepository
	customerRepo infrastructure.CustomerRepository
}

//NewLoginHandler construct handler
func NewLoginHandler(loginRepo infrastructure.LoginRepository, accountRepo infrastructure.AccountRepository, customerRepo infrastructure.CustomerRepository) LoginHandler {
	return LoginHandler{
		repository:   loginRepo,
		accountRepo:  accountRepo,
		customerRepo: customerRepo,
	}
}

//LoginHandler method for login
func (handler *LoginHandler) LoginHandler(c *gin.Context) {
	request := Request{}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		errRequest := helper.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(errRequest, false))
		return
	}

	//Get user by email
	account, err := handler.accountRepo.GetAccountByEmail(ctx, request.Data.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
		return
	}

	//compare password
	byteDBPass := []byte(account.Password)
	byteReqPass := []byte(request.Data.Password)

	if error := bcrypt.CompareHashAndPassword(byteDBPass, byteReqPass); error != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage("Password Anda Salah", false))
		return
	}

	// Create the Claims
	claims := helper.CreateClaims(account.ID, account.Email, account.Role, helper.ExpiredAt)
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims) //Create Token
	signed, err := token.SignedString([]byte("secret"))

	//Login handler
	_, errCreateAccessToken := handler.repository.Login(ctx, RequestMapper(account.ID, signed), account.ID)
	if errCreateAccessToken != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(errCreateAccessToken.Error(), false))
		return
	}

	accountID := strconv.FormatUint(account.ID, 10)

	// Get Account By ID
	accountData, errGetAccount := handler.accountRepo.GetAccountByID(ctx, accountID)
	if errGetAccount != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(errGetAccount.Error(), false))
		return
	}

	// Get Customer By ID
	customerData, errGetAccount := handler.customerRepo.GetCustomerByID(ctx, accountID)
	if errGetAccount != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(errGetAccount.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(accountData, customerData), "Login Berhasil", true))
}
