package entities

import (
	base "altastore-api/domain/persistence"
)

//Category struct models
type Category struct {
	base.ModelSoftDelete
	Name string `gorm:"column:name" json:"name"`
}
