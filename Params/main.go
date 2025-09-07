package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/user/:name", func(c *fiber.Ctx) error {
		fmt.Println(c.AllParams())
		return c.JSON(c.AllParams()) // "{"name": "fenny"}"
	})

	app.Get("/user/*", func(c *fiber.Ctx) error {
		return c.JSON(c.AllParams()) // "{"*1": "fenny/123"}"

		// ...
	})

	app.Listen(":3000")
}
