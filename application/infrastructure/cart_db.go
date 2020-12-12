package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

//CartRepository repository
type CartRepository interface {
	CreateCart(context.Context, domain.Cart, []domain.CartDetail) (domain.Cart, error)
	GetListCart(context.Context, string) []domain.Cart
	DeleteCartDetail(context.Context, string) error
}
