package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.SendFile("./public/inedx.html")
	})
	app.Listen(":3000")
}
