package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

//CustomerRepository customer respository
type CustomerRepository interface {
	GetCustomerByID(context.Context, string) (domain.Customer, error)
}
