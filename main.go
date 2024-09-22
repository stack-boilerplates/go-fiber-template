package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/vserpa/go-fiber-template/database"
	"github.com/vserpa/go-fiber-template/router"

	_ "github.com/lib/pq"
)

func setupApp() *fiber.App {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(logger.New())

	router.SetupApiRoutes(app)

	app.Static("/home", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("the value is: " + c.Params("value"))
	})

	return app
}

func main() {
	app := setupApp()
	app.Listen(":3000")
}
