package infrastructure

import (
	domain "altastore-api/domain/entities"
	"context"
)

//CartRepository repository
type CartRepository interface {
	CreateCart(context.Context, domain.Cart) (domain.Cart, error)
	CreateCartDetail(context.Context, domain.CartDetail) (domain.Cart, error)
	GetCartByCustomerID(context.Context, string) (domain.Cart, error)
	GetCartByCartID(context.Context, string) (domain.Cart, error)
	GetListCartDetail(context.Context, string, domain.FilterRequest) []domain.Cart
	DeleteCartDetailByID(context.Context, string) error
	UpdateCartByCartID(context.Context, string, domain.Cart) (domain.Cart, error)
}
