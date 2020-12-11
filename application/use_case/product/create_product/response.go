package create_product

import (
	"altastore-api/application/helper"
	domain "altastore-api/domain/entities"
	"time"
)

type (
	CreateProductResponse struct {
		helper.BaseResponse
		Data CreateProductResponseData `json:"data"`
	}
)

type CreateProductResponseData struct {
	ID          uint64    `json:"id"`
	CategoryID  string    `json:"category_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Stock       uint      `json:"stock"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func SetResponse(data CreateProductResponseData, message string, success bool) CreateProductResponse {
	return CreateProductResponse{
		BaseResponse: helper.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: data,
	}
}

func ResponseMapper(domain domain.Product) CreateProductResponseData {
	return CreateProductResponseData{
		ID:          domain.ModelSoftDelete.ID,
		CategoryID:  domain.CategoryID,
		ProductName: domain.ProductName,
		Description: domain.Description,
		Stock:       domain.Stock,
		Price:       domain.Price,
		CreatedAt:   domain.ModelSoftDelete.CreatedAt,
		UpdatedAt:   domain.ModelSoftDelete.UpdatedAt,
		DeletedAt:   domain.ModelSoftDelete.UpdatedAt,
	}
}
