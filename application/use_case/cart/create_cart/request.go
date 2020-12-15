package create_cart

import (
	domain "altastore-api/domain/entities"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

type (
	//CreateCartRequest create cart request struct
	CreateCartRequest struct {
		Data struct {
			ProductID int `json:"product_id" validate:"required"`
			Quantity  int `json:"quantity" validate:"required"`
		}
	}
)

//ValidateRequest create cart request
func ValidateRequest(req *CreateCartRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

//RequestMapper mapper from request to domain
func RequestMapper(accountID string) domain.Cart {
	customerID, _ := strconv.ParseUint(accountID, 0, 64)
	return domain.Cart{
		CustomerID:       customerID,
		Quantity:         0,
		CartTotalAmmount: 0,
	}
}
