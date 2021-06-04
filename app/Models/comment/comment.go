package comment

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Email string `gorm:"column:email"`
	Name uint `gorm:"column:name"`
	PostsId uint `gorm:"column:posts_id"`
	Content string
	Status string
}