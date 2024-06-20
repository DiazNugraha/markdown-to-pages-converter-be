package routes

import (
	"markdown-to-pages-converter/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
	route.Post("/generate", controllers.GeneratePages)
}
