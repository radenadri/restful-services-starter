package routes

import (
	"boilerplate/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router) {
	app.Get("/users/:id", controllers.ShowUser)
}
