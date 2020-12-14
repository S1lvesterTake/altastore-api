package create_cart

import "gopkg.in/go-playground/validator.v9"

type (
	CreateCartRequest struct {
		Data struct {
			ProductID string `json:"product_id" validate:"required"`
			Quantity  string `json:"quantity" validate:"required"`
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
