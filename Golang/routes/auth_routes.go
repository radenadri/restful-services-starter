package routes

import (
	"boilerplate/app/controllers"
	"boilerplate/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router) {
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)
	app.Post("/refresh", controllers.RefreshToken)
	app.Get("/me", middlewares.EnableJWT(), controllers.Profile)
	app.Put("/me", middlewares.EnableJWT(), controllers.UpdateProfile)
}
