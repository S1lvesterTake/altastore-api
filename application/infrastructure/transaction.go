package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

//TransctionRepository repository
type TransctionRepository interface {
	CreateTransaction(context.Context, domain.Transaction, []domain.TransactionDetail) (domain.Transaction, error)
	GetListTransaction(context.Context, string) []domain.Transaction
}
