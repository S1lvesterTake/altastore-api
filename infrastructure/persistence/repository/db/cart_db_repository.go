package db

import (
	"altastore-api/application/infrastructure"
	"context"
	"errors"
	"fmt"
	"strconv"

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

func (c *cartRepository) CreateCart(ctx context.Context, quantity, productID int, customerID string) (domain.Cart, error) {

	tx := c.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var cartID uint64
	custID, _ := strconv.ParseUint(customerID, 0, 64)

	product := domain.Product{}
	cart := domain.Cart{
		CustomerID: custID,
	} // insert value

	existingCart, err := c.GetCartByCustomerID(ctx, customerID)
	if err != nil {
		if err.Error() == "record not found" {
			//create new cart
			err := tx.Create(&cart).Error
			if err != nil {
				tx.Rollback()
				return cart, err
			}
		}
		// return cart, err
	}

	if existingCart.ID != 0 {
		cartID = existingCart.ID
	} else {
		cartID = cart.ID
	}

	//find product by product ID
	prodID := fmt.Sprint(productID)
	if c.DB.First(&product, "id = ?", prodID).RecordNotFound() {
		return cart, errors.New("Produk dengan id " + prodID + " tidak ditemukan")
	}

	//create cart detail
	cartDetail := domain.CartDetail{
		CartID:           cartID,
		ProductID:        product.ID,
		Quantity:         quantity,
		CartDetailAmount: float64(quantity) * float64(product.Price),
	}
	err = tx.Create(&cartDetail).Error
	if err != nil {
		tx.Rollback()
		return cart, err
	}

	if err := tx.Commit().Error; err != nil {
		return cart, err
	}

	if existingCart.ID != 0 {
		return existingCart, nil
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

	err := c.DB.Where("customer_id = ?", customerID).First(&cartData).Error
	if err != nil {
		return cartData, err
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

func (c *cartRepository) GetAllCartDetailByCartID(ctx context.Context, cartID string) ([]domain.CartDetail, error) {
	var cartDetails []domain.CartDetail

	if c.DB.Where("cart_id = ?", cartID).Preload("Product").Find(&cartDetails).RecordNotFound() {
		return cartDetails, errors.New("Cart Detail dengan CartID " + cartID + " tidak ditemukan")
	}
	return cartDetails, nil
}
