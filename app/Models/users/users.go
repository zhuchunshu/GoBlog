package users

import (
	"database/sql"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName string
	PassWord string
	Email string
	EmailVerTime sql.NullTime
	Shenfen string
}