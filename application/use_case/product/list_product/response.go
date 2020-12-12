package list_product

import (
	"altastore-api/application/helper"
	domain "altastore-api/domain/entities"
)

type (
	ProductListResponse struct {
		helper.BaseResponse
		Pagination domain.FilterResponse `json:"pagination"`
		Data       []domain.Product      `json:"product"`
	}
)

// SetResponse set product list response
func SetResponse(domain []domain.Product, pagination domain.FilterResponse, message string, success bool) ProductListResponse {
	return ProductListResponse{
		BaseResponse: helper.BaseResponse{
			Success: success,
			Message: message,
		},
		Pagination: pagination,
		Data:       domain,
	}
}
