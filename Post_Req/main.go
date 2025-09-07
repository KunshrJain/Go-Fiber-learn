package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func main() {
	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		var u User
		if err := c.BodyParser(&u); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		fmt.Println("On no")

		// pretend to save to database
		return c.JSON(fiber.Map{
			"status":   "ok",
			"username": u.Username,
			"age":      u.Age,
		})
	})

	app.Listen(":3000")
}
