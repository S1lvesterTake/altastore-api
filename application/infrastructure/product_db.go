package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

//ProductRepository repository
type ProductRepository interface {
	CreateProduct(context.Context, domain.Product) (domain.Product, error)
	CreateCategory(context.Context, string) (domain.Category, error)
	GetListProductByCategory(context.Context, string) []domain.Product
}
