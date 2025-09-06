package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public") // Serve static files from the "public" directory

	// Route to handle root path

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Guys!")
	})

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})
	app.Get("/:value", func(c *fiber.Ctx) error {
		fmt.Println("value: " + c.Params("value"))
		return c.SendString("value: " + c.Params("value"))
	})
	// GET http://localhost:3000/api/user/john

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => API path: user/john
	})

	app.Listen(":3000")
}
