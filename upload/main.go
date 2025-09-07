package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// func handler2(app *fiber.App) {
// 	app.Get("/download", func(c *fiber.Ctx) error {
// 		return c.SendString("File downloaded successfully")
// 	})
// }

func handler1(app *fiber.App) {
	app.Get("/download", func(c *fiber.Ctx) error {
		defer c.Download("./public/sample.pdf")
		return nil
	})
}
func main() {
	app := fiber.New()
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, _ := c.FormFile("document")
		return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
	})
	app.Listen(":3000")
}
