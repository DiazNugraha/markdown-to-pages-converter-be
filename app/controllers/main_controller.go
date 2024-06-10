package controllers

import (
	"log"
	"markdown-to-pages-converter/app/services"

	"github.com/gofiber/fiber/v2"
)

func GeneratePages(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["file"]
	if len(files) == 0 {
		log.Println("No file uploaded")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "No file uploaded",
		})
	}

	if err := services.MainService(ctx, files); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "File buffer error: " + err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}
