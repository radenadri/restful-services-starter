package middlewares

import (
	"boilerplate/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Limiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(utils.Response{
				Status:  false,
				Message: "Too many requests, please try again later",
				Data:    nil,
			})
		},
	})
}
