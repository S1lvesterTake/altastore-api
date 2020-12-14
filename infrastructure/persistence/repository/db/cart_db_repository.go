package db

import (
	"altastore-api/application/infrastructure"
	"context"
	"errors"

	domain "altastore-api/domain/entities"

	"github.com/jinzhu/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

//NewCartRepository construct cart repository
func NewCartRepository(DB *gorm.DB) infrastructure.CartRepository {
	return &cartRepository{
		DB: DB,
	}
}

func (c *cartRepository) CreateCart(ctx context.Context, cart domain.Cart) (domain.Cart, error) {
	if err := c.DB.Create(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (c *cartRepository) CreateCartDetail(ctx context.Context, cartDetail domain.CartDetail) (domain.Cart, error) {
	var cart domain.Cart
	if err := c.DB.Create(&cartDetail).Error; err != nil {
		return cart, err
	}
	return cart, nil
}
func (c *cartRepository) GetListCartDetail(ctx context.Context, cartID string, pagination domain.FilterRequest) []domain.Cart {

	return nil
}

func (c *cartRepository) DeleteCartDetailByID(ctx context.Context, cartDetailID string) error {
	err := c.DB.Delete(&domain.CartDetail{}, "id = ? ", cartDetailID).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) GetCartByCustomerID(ctx context.Context, customerID string) (domain.Cart, error) {
	cartData := domain.Cart{}

	if c.DB.Where("customer_id = ?", customerID).First(&cartData).RecordNotFound() {
		return cartData, errors.New("Cart dengan CustomerID " + customerID + " tidak ditemukan")
	}
	return cartData, nil
}

func (c *cartRepository) GetCartByCartID(ctx context.Context, cartID string) (domain.Cart, error) {
	cart := domain.Cart{}

	if c.DB.Where("id = ?", cartID).Preload("CartDetail").Preload("Product").First(&cart).RecordNotFound() {
		return cart, errors.New("Cart dengan CartID " + cartID + " tidak ditemukan")
	}
	return cart, nil
}

func (c *cartRepository) UpdateCartByCartID(ctx context.Context, cartID string, cart domain.Cart) (domain.Cart, error) {
	if err := c.DB.Model(&domain.Cart{}).Where("id = ?", cartID).Update(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}
