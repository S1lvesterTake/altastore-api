package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

//AccountRepository account respository
type AccountRepository interface {
	GetAccountByEmail(context.Context, string) (domain.Account, error)
	GetAccountByID(context.Context, string) (domain.Account, error)
}
