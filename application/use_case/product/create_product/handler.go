package create_product

import (
	"altastore-api/application/helper"
	"altastore-api/application/infrastructure"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

//CreateProductHandler handler struct
type CreateProductHandler struct {
	request    infrastructure.Request
	repository infrastructure.ProductRepository
}

//NewCreateProductHandler construct handler
func NewCreateProductHandler(request infrastructure.Request, repo infrastructure.ProductRepository) CreateProductHandler {
	return CreateProductHandler{
		request:    request,
		repository: repo,
	}
}

//CreateProductHandler handler create product
func (handler *CreateProductHandler) CreateProductHandler(c *gin.Context) {
	request := CreateProductRequest{}
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	role, _ := c.Get("Role")
	if role.(string) != helper.ADMIN {
		c.JSON(http.StatusUnauthorized, response.SetMessage("Anda bukan admin", false))
		return
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

	_, err := handler.repository.GetCategoryByID(ctx, request.Data.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
		return
	}

	product, err := handler.repository.CreateProduct(ctx, RequestMapper(request))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(product), "Produk berhasil dibuat", true))
}
