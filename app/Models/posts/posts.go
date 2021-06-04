package posts

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title string
	Content string
	UserId string `gorm:"column:user_id"`
	Status string
	Type string
}