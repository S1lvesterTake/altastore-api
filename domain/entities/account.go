package entities

import (
	base "altastore-api/domain/persistence"
	"time"
)

// Account struct models
type Account struct {
	base.Model
	Role        string      `gorm:"column:role" json:"role"`
	Email       string      `gorm:"column:email" json:"email"`
	Password    string      `gorm:"column:password" json:"password"`
	LastLogin   time.Time   `gorm:"column:last_login" json:"last_login"`
	AccessToken AccessToken `json:"access_token"`
}
