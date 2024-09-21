package main

import "github.com/gofiber/fiber/v2"

func setupApp() *fiber.App {
	app := fiber.New()

	app.Static("/home", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})

	return app
}

func main() {
	app := setupApp()
	app.Listen(":3000")
}
