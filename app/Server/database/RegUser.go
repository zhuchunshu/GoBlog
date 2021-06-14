package database

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/zhuchunshu/GoBlog/app/Models/users"
	"time"
)

func RegUser()  {
	db:=DBConn
	var count int64
	db.Table("users").Count(&count)
	if count==0 {
		pwd, _ :=gmd5.EncryptString("123456")
		db.Model(&users.Users{}).Create([]map[string]interface{}{
			{"username": "admin", "password": pwd,"email":"laravel@88.com","shenfen":"admin","created_at":time.Now(),"updated_at":time.Now()},
		})
	}
}