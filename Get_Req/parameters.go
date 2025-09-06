package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Parameters() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Guys!")
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		fmt.Println("value: " + c.Params("value"))
		return c.SendString("value: " + c.Params("value"))
	})

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})

	app.Listen(":3000")
}
