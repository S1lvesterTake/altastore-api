package create_cart

import (
	"altastore-api/application/helper"
	"altastore-api/application/infrastructure"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

/*
only Cutomer role can create cart
functionalitiy :
if user have cart
true -> insert into table cart detail
false -> insert into table cart and cart detail
*/

//CreateCartHandler handler struct
type CreateCartHandler struct {
	request     infrastructure.Request
	repository  infrastructure.CartRepository
	productRepo infrastructure.ProductRepository
}

//NewCreateCartHandler construct hanbler depedency
func NewCreateCartHandler(request infrastructure.Request, repo infrastructure.CartRepository, productRepo infrastructure.ProductRepository) CreateCartHandler {
	return CreateCartHandler{
		request:     request,
		repository:  repo,
		productRepo: productRepo,
	}
}

//CreateCartHandler create cart handler implementation
func (handler *CreateCartHandler) CreateCartHandler(c *gin.Context) {
	//check  request
	request := CreateCartRequest{}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	accountID, _ := c.Get("AccountId")
	accID := fmt.Sprint(accountID.(uint64))
	role, _ := c.Get("Role")
	if role.(string) != helper.CUSTOMER {
		c.JSON(http.StatusUnauthorized, response.SetMessage("Anda bukan admin", false))
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	//validate request
	if ok, err := ValidateRequest(&request); !ok {
		errRequest := helper.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(errRequest, false))
		return
	}

	cart, err := handler.repository.CreateCart(ctx, request.Data.Quantity, request.Data.ProductID, accID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
		return
	}

	//find cart detail by cart id
	cartID := fmt.Sprint(cart.ID)
	cartDetails, err := handler.repository.GetAllCartDetailByCartID(ctx, cartID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
		return
	}
	var cartDetailTotalQty int
	//find total quantity
	for _, cartItem := range cartDetails {
		cartDetailTotalQty += cartItem.Quantity
	}

	//update cart quantity
	cart.Quantity = cartDetailTotalQty
	_, err = handler.repository.UpdateCartByCartID(ctx, cartID, cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetReponse(ResponseMapper(cart, cartDetails), "Berhasil menambahkan keranjang", true))
}
