package entities

import (
	base "altastore-api/domain/persistence"
)

//Customer struct models
type Customer struct {
	base.ModelSoftDelete
	Name        string `gorm:"column:name" json:"name"`
	PhoneNumber string `gorm:"column:phone_number" json:"email"`
	Address     string `gorm:"column:address" json:"address"`
	Gender      string `gorm:"column:gender" json:"gender"`
}
