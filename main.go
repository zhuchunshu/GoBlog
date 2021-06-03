package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zhuchunshu/GoBlog/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "FecApi",
	})
	app.Use(logger.New())

	// 初始化数据库
	//initDatabase()

	// 注册路由
	routes.InitRouters(app)

	// 设置端口
	//Serverport := helpers.JsonDecode(ServerConfig, "port")
	server := app.Listen(":3000")
	log.Fatal(server)
}
