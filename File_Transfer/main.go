package main

import (
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

	handler1(app)

	app.Listen(":3000")
}
