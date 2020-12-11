package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

//ProductRepository repository
type ProductRepository interface {
	CreateProduct(context.Context, domain.Product) (domain.Product, error)
	GetListProductByCategory(context.Context, domain.PaginationRequest) []domain.Product
	GetCategoryByID(context.Context, string) (domain.Category, error)
}
