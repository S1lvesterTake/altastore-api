package entities

import (
	base "altastore-api/domain/persistence"
)

//Product struct models
type (
	Product struct {
		base.ModelSoftDelete
		CategoryID  string `gorm:"column:category_id" json:"category_id"`
		ProductName string `gorm:"column:product_name" json:"product_name"`
		Description string `gorm:"column:description" json:"description"`
		Stock       uint   `gorm:"column:stock" json:"stock"`
		Price       int64  `gorm:"column:price" json:"price"`
	}
)
