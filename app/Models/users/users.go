package users

import (
	"database/sql"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName string  `gorm:"column:username"`
	PassWord string `gorm:"column:password"`
	Email string
	EmailVerTime sql.NullTime `gorm:"column:email_ver_time"`
	Shenfen string
}