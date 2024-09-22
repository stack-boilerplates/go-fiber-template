package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/vserpa/go-fiber-template/handler"
	"github.com/vserpa/go-fiber-template/middleware"
)

// SetupApiRoutes func
func SetupApiRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New(), middleware.AuthReqMiddleware())

	// routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)
}
