package create_cart

import (
	"altastore-api/application/infrastructure"

	"github.com/gin-gonic/gin"
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
	request    infrastructure.Request
	repository infrastructure.CartRepository
}

//NewCreateCartHandler construct hanbler depedency
func NewCreateCartHandler(request infrastructure.Request, repo infrastructure.CartRepository) CreateCartHandler {
	return CreateCartHandler{
		request:    request,
		repository: repo,
	}
}

func (handler *CreateCartHandler) CreateCartHandler(c *gin.Context) {

}
