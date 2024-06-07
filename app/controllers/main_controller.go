package controllers

import (
	"fmt"
	"log"

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

	for _, file := range files {
		err := ctx.SaveFile(file, fmt.Sprintf("./public/%s", file.Filename))
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("success save")
	}

	return ctx.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}
