package routes

import (
	"boilerplate/app/controllers"
	"boilerplate/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PostRoute(app fiber.Router) {
	app.Get("/posts", controllers.IndexPost)
	app.Get("/posts/:id", controllers.ShowPost)
	app.Post("/posts", middlewares.EnableJWT(), controllers.StorePost)
	app.Put("/posts/:id", middlewares.EnableJWT(), controllers.UpdatePost)
	app.Delete("/posts/:id", middlewares.EnableJWT(), controllers.DestroyPost)
}
