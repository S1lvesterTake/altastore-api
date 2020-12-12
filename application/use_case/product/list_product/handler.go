package list_product

import (
	"altastore-api/application/infrastructure"
	domain "altastore-api/domain/entities"
	"context"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ListProductHandler list product handler
type ListProductHandler struct {
	request    infrastructure.Request
	repository infrastructure.ProductRepository
}

//NewListProductHandler construct new List  product handler
func NewListProductHandler(req infrastructure.Request, repo infrastructure.ProductRepository) ListProductHandler {
	return ListProductHandler{
		request:    req,
		repository: repo,
	}
}

//ListProductHandler implementtation
func (handler *ListProductHandler) ListProductHandler(c *gin.Context) {

	var (
		firstRowOnPage int
		lastRowOnPage  int
		page, _        = strconv.Atoi(c.Query("paging[page]"))
		limit, _       = strconv.Atoi(c.DefaultQuery("paging[limit]", "10"))
	)

	lastRowOnPage = page * limit
	if page == 1 {
		firstRowOnPage = page
	} else {
		firstRowOnPage = lastRowOnPage - limit + 1
	}

	filter := domain.FilterRequest{
		Filter: c.Query("filter[category]"),
		Page:   c.Query("paging[page]"),
		Limit:  c.DefaultQuery("paging[limit]", "10"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	productsData, count := handler.repository.GetListProduct(ctx, filter)

	if page == 0 {
		firstRowOnPage = 1
		lastRowOnPage = len(productsData)
	}
	paginationResponse := domain.FilterResponse{
		CurrentPage:    page,
		PageCount:      int(math.Ceil(float64(count) / float64(limit))),
		PageSize:       limit,
		RowCount:       count,
		FirstRowOnPage: firstRowOnPage,
		LastRowOnPage:  lastRowOnPage,
	}

	c.JSON(http.StatusOK, SetResponse(productsData, paginationResponse, "List Product berhasil dimuat ", true))
	return
}
