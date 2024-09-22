package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"

	"github.com/vserpa/go-fiber-template/config"
)

func AuthReqMiddleware() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	})
}
