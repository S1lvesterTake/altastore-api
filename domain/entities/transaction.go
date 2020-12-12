package entities

import (
	base "altastore-api/domain/persistence"
)

//Transaction struct models
type Transaction struct {
	base.Model
	CustomerID        uint64              `gorm:"column:customer_id" json:"customer_id"`
	TransactionNumber uint64              `gorm:"column:transaction_number" json:"transaction_number"`
	CodeID            string              `gorm:"column:code_id" json:"code_id"`
	TotalAmount       int64               `gorm:"column:total_amount" json:"total_amount"`
	TransactionDetail []TransactionDetail `json:"transaction_details"`
}
