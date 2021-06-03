package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zhuchunshu/FecApi-Api/app/server/config"
	"github.com/zhuchunshu/FecApi-Api/helpers"
	"github.com/zhuchunshu/GoBlog/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

// 服务配置
var ServerConfig string = config.GetServerConfig()

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "FecApi",
	})
	app.Use(logger.New())

	// 初始化数据库
	initDatabase()

	// 注册路由
	routes.InitRouters(app)

	// 设置端口
	Serverport := helpers.JsonDecode(ServerConfig, "port")
	server := app.Listen(":" + Serverport)
	log.Fatal(server)
}
