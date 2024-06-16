package middlewares

import (
	"boilerplate/app/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: config.CORS_ALLOWED_ORIGINS,
		AllowHeaders: "Origin, Content-Type, Accept",
	})
}
