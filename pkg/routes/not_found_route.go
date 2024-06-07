package routes

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(app *fiber.App) {
	app.Use(func(ctx *fiber.Ctx) error {
		code := fiber.StatusNotFound
		return ctx.Status(code).JSON(fiber.Map{
			"error": true,
			"msg":   "sorry, endpoint not found",
		})
	})
}
