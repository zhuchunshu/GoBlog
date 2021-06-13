package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitRouters(app *fiber.App) {
	Web(app)
	//RegisterWebRoutes(app)
	//api.InitRouters(app)
}
