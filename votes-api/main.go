package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/api/v1/votes", func(c *fiber.Ctx) error {
		return c.SendString("Votes API")
	})

	app.Listen(":3000")
}
