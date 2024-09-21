package main

import "github.com/gofiber/fiber/v2"

func setupApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}

func main() {
	app := setupApp()
	app.Listen(":3000")
}
