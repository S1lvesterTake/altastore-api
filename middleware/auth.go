package auth

import (
	domain "altastore-api/domain/entities"
	"altastore-api/infrastructure/persistence/repository/db"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

/*
Token JWT claims struct
*/
type Token struct {
	AccountID uint64
	Email     string
	Role      string
	jwt.StandardClaims
}

// AuthenticationRequired for auth
func AuthenticationRequired() gin.HandlerFunc {

	var accToken domain.AccessToken

	return func(c *gin.Context) {
		notAuth := []string{
			"/healthcheck",
			"/api/v1/login",
		} //List of endpoints that doesn't require auth
		requestPath := c.Request.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath || strings.HasPrefix(requestPath, value) {
				c.Next()
				return
			}
		}

		tokenHeader := c.Request.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			c.JSON(http.StatusForbidden, response.SetMessage("Missing Auth Token", false))
			c.Abort()
			return
		}
		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			c.JSON(http.StatusForbidden, response.SetMessage("Invalid/Malformed auth token", false))
			c.Abort()
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		//handle if user by email status 0 abort next handler
		account := domain.Account{}
		if err := db.GetDB().First(&account, tk.AccountID).Error; err != nil {
			c.JSON(http.StatusForbidden, response.SetMessage("User is not in DB.", false))
			c.Abort()
			return
		}

		if err != nil { //Malformed token, returns with http code 403 as usual
			c.JSON(http.StatusForbidden, response.SetMessage(fmt.Sprintf("Malformed authentication token: %s", err.Error()), false))
			c.Abort()
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			c.JSON(http.StatusForbidden, response.SetMessage("Token is not valid.", false))
			c.Abort()
			return
		}

		// Check token
		if db.GetDB().Table("access_tokens").Select("token").Where(&domain.AccessToken{Token: tokenPart}).First(&accToken).RecordNotFound() {
			c.JSON(http.StatusForbidden, response.SetMessage("Session login Anda berubah. Silakan lakukan login ulang", false))
			c.Abort()
			return
		}

		c.Set("AccountId", tk.AccountID)
		c.Set("Email", tk.Email)
		c.Set("Role", tk.Role)
		c.Next() //proceed in the core chain!
	}
}
