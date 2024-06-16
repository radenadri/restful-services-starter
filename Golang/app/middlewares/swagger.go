package middlewares

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func Swagger() fiber.Handler {
	return swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/api-docs",
		Title:    "Go Fiber boilerplate API",
		CacheAge: 0,
	})
}
