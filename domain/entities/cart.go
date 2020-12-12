package entities

import (
	base "altastore-api/domain/persistence"
)

//Cart struct models
type Cart struct {
	base.Model
	CustomerID       uint64       `gorm:"column:customer_id" json:"customer_id"`
	CodeID           string       `gorm:"column:code_id" json:"code_id"`
	Quantity         int          `gorm:"column:quantity" json:"quantity"`
	CartTotalAmmount int64        `gorm:"column:cart_total_amount" json:"cart_total_amount"`
	CartDetail       []CartDetail `json:"cart_details"`
}
