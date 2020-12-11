package create_product

import (
	domain "altastore-api/domain/entities"

	"gopkg.in/go-playground/validator.v9"
)

type (
	//Request login request struct
	CreateProductRequest struct {
		Data struct {
			CategoryID  string `json:"category_id" validate:"required"`
			Name        string `json:"name" validate:"required"`
			Description string `json:"description"`
			Stock       uint   `json:"stock"`
			Price       int64  `json:"price"`
		}
	}
)

//ValidateRequest validate request create product
func ValidateRequest(req *CreateProductRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

//RequestMapper mapper request to domain
func RequestMapper(req CreateProductRequest) domain.Product {
	return domain.Product{
		CategoryID:  req.Data.CategoryID,
		ProductName: req.Data.Name,
		Description: req.Data.Description,
		Stock:       req.Data.Stock,
		Price:       req.Data.Price,
	}
}
