package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// GET endpoint returning JSON
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":  "Kunsh",
			"role":  "Engineering Student",
			"stack": "Go + Fiber",
		})
	})

	app.Listen(":3000")
}
