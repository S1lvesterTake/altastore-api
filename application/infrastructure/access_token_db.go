package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

// LoginRepository repository contract
type LoginRepository interface {
	Login(context.Context, domain.AccessToken, uint64) (domain.AccessToken, error)
	// Register(context.Context)
}
