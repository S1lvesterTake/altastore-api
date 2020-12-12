package entities

import (
	base "altastore-api/domain/persistence"
)

//CartDetail struct models
type CartDetail struct {
	base.ModelSoftDelete
	CartID           uint64 `gorm:"column:cart_id" json:"cart_id"`
	ProductID        uint64 `gorm:"column:product_id" json:"product_id"`
	Quantity         int    `gorm:"column:code_id" json:"code_id"`
	CartDetailAmount int64  `gorm:"column:cart_detail_amount" json:"cart_detail_amount"`
}
