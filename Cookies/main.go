package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Cookie struct {
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	Path        string    `json:"path"`
	Domain      string    `json:"domain"`
	MaxAge      int       `json:"max_age"`
	Expires     time.Time `json:"expires"`
	Secure      bool      `json:"secure"`
	HTTPOnly    bool      `json:"http_only"`
	SameSite    string    `json:"same_site"`
	SessionOnly bool      `json:"session_only"`
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		// Create cookie
		// cookie := new(fiber.Cookie)
		// cookie.Name = "john"
		// cookie.Value = "doe"
		// cookie.Expires = time.Now().Add(24 * time.Hour)

		// cookie.SameSite = "Lax"
		// cookie.Path = "/"
		// cookieTheme := new(fiber.Cookie)
		// cookieTheme.Name = "john"
		// cookieTheme.Value = "light"
		// cookieTheme.MaxAge = 100

		// c.Cookie(cookie)
		// c.Cookie(cookieTheme)
		// ...

		// c.ClearCookie()

		return nil
	})
	app.Listen(":3000")
}
