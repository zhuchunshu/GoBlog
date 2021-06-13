package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/zhuchunshu/FecApi-Api/app/server/config"
	"github.com/zhuchunshu/FecApi-Api/helpers"
	"github.com/zhuchunshu/GoBlog/app/Models/comment"
	"github.com/zhuchunshu/GoBlog/app/Models/posts"
	"github.com/zhuchunshu/GoBlog/app/Models/users"
	"github.com/zhuchunshu/GoBlog/app/Server/database"
	"github.com/zhuchunshu/GoBlog/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

// 服务配置
var ServerConfig = config.GetServerConfig()

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		ServerHeader:  "FecApi",
		Views: engine,
	})
	app.Use(logger.New())

	// 初始化静态资源目录
	initPublicPath(app, helpers.JsonDecode(ServerConfig, "public"))
	// 初始化数据库
	initDatabase()

	// 数据库迁移
	_ = database.DBConn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&users.Users{}, &posts.Posts{}, &comment.Comment{})
	// 注册路由
	routes.InitRouters(app)

	// 设置端口
	Serverport := helpers.JsonDecode(ServerConfig, "port")
	server := app.Listen(":" + Serverport)
	log.Fatal(server)
}
