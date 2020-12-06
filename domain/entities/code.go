package entities

import "time"

//Code struct models
type Code struct {
	ID        string    `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
