package create_cart

import (
	"altastore-api/application/helper"
	domain "altastore-api/domain/entities"
)

type (
	//CreateCartResponse .
	CreateCartResponse struct {
		helper.BaseResponse
		Data CreateCartResponseData `json:"data"`
	}
	//CreateCartResponseData .
	CreateCartResponseData struct {
		ID         uint64         `json:"id"`
		CustomerID uint64         `json:"customer_id"`
		Quantity   int            `json:"quantity"`
		CartItems  []CartItemData `json:"cart_items"`
	}
	//CartItemData .
	CartItemData struct {
		ID             uint64  `json:"id"`
		CartID         uint64  `json:"cart"`
		ProductID      uint64  `json:"product_id"`
		Quantity       int     `json:"quantity"`
		CartItemAmount float64 `json:"cart_item_amount"`
		Product        Product `json:"product"`
	}
	//Product .
	Product struct {
		CategoryID  string `json:"category_id"`
		ProductName string `json:"product_name"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
	}
)

//ResponseMapper create cart response mapper
func ResponseMapper(cartData domain.Cart, cartItemsData []domain.CartDetail) CreateCartResponseData {
	var cartItems []CartItemData

	if len(cartItemsData) != 0 {
		for _, item := range cartItemsData {
			cartItems = append(cartItems, CartItemData{
				ID:             item.ID,
				CartID:         item.CartID,
				ProductID:      item.ProductID,
				Quantity:       item.Quantity,
				CartItemAmount: item.CartDetailAmount,
				Product: Product{
					CategoryID:  item.Product.CategoryID,
					ProductName: item.Product.ProductName,
					Description: item.Product.Description,
					Price:       item.Product.Price,
				},
			})
		}
	}

	return CreateCartResponseData{
		ID:         cartData.ID,
		CustomerID: cartData.CustomerID,
		Quantity:   cartData.Quantity,
		CartItems:  cartItems,
	}
}

//SetReponse create cart set response
func SetReponse(data CreateCartResponseData, message string, isSuccess bool) CreateCartResponse {
	return CreateCartResponse{
		BaseResponse: helper.BaseResponse{
			Message: message,
			Success: isSuccess,
		},
		Data: data,
	}
}
