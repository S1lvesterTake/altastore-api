package login

import (
	base "altastore-api/application/helper"
	domain "altastore-api/domain/entities"
	"time"
)

type (
	//Response login response
	Response struct {
		base.BaseResponse
		Data ResponseData `json:"data"`
	}

	//ResponseData login response data
	ResponseData struct {
		AccountID   uint64    `json:"account_id"`
		Email       string    `json:"email"`
		Name        string    `json:"name"`
		Role        string    `json:"role"`
		Address     string    `json:"address"`
		PhoneNumber string    `json:"phone_number"`
		Token       string    `json:"token"`
		ExpiredAt   int64     `json:"expired_at"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

//SetResponse for login
func SetResponse(domain ResponseData, message string, success bool) Response {
	return Response{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

//ResponseMapper login response mapper from db
func ResponseMapper(accountData domain.Account, customerData domain.Customer) ResponseData {

	return ResponseData{
		AccountID:   accountData.ID,
		Email:       accountData.Email,
		Name:        customerData.Name,
		Role:        accountData.Role,
		Address:     customerData.Address,
		PhoneNumber: customerData.PhoneNumber,
		Token:       accountData.AccessToken.Token,
		ExpiredAt:   accountData.AccessToken.ExpiredAt,
		CreatedAt:   accountData.CreatedAt,
		UpdatedAt:   accountData.UpdatedAt,
	}
}
