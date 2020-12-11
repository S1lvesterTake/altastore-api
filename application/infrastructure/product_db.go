package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

//ProductRepository repository
type ProductRepository interface {
	CreateProduct(context.Context, domain.Product) (domain.Product, error)
	GetListProduct(context.Context, domain.FilterRequest) []domain.Product
	GetCategoryByID(context.Context, string) (domain.Category, error)
}
