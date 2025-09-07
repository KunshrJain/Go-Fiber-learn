package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/yt", func(c *fiber.Ctx) error {
		return c.Redirect("https://www.youtube.com/")
	})

	app.Get("/:page", func(c *fiber.Ctx) error {
		page := c.Params("page")
		return c.SendFile("./public/" + page + ".html")
	})

	app.Listen(":3000")
}
