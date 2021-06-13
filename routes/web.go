package routes

import "github.com/gofiber/fiber/v2"

func Web(route *fiber.App)  {
	route.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
}