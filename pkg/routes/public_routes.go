package routes

import (
	"markdown-to-pages-converter/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/generate", controllers.GeneratePages)
}
