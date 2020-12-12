package entities

import (
	base "altastore-api/domain/persistence"
)

//TransactionDetail struct models
type TransactionDetail struct {
	base.Model
	TransactionID uint64 `gorm:"column:transaction_id" json:"transaction_id"`
	ProductID     uint64 `gorm:"column:product_id" json:"product_id"`
	Quantity      string `gorm:"column:quantity" json:"quantity"`
	DetailAmount  int64  `gorm:"column:detail_amount" json:"detail_amount"`
}
