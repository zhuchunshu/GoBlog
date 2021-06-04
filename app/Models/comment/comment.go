package comment

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId uint `gorm:"column:user_id"`
	PostsId uint `gorm:"column:posts_id"`
	Content string
	Status string
}