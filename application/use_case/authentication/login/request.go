package login

import (
	"altastore-api/application/helper"
	domain "altastore-api/domain/entities"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	//Request login request struct
	Request struct {
		Data struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required"`
		}
	}
)

//ValidateRequest login request
func ValidateRequest(req *Request) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

// RequestMapper login request mapper
func RequestMapper(userID uint64, signed string) domain.AccessToken {
	return domain.AccessToken{
		AccountID: userID,
		Token:     signed,
		ExpiredAt: helper.GetExpiryTime(time.Duration(helper.ExpiredAt)),
	}
}
