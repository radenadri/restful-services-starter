package main

import (
	"boilerplate/app/config"
	pg "boilerplate/app/db"
	"boilerplate/app/middlewares"
	"boilerplate/routes"

	"github.com/gofiber/fiber/v2"
)

// @title           Go Fiber + Gorm Boilerplate
// @version         1.0
// @description     This is a simple CRUD API application made with Golang and documented with Swagger
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9000
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// Init database
	pg.Init()

	// Migrate database
	pg.Migrate()

	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	// Enable CORS
	app.Use(middlewares.Cors())

	// Enable limiter
	app.Use(middlewares.Limiter())

	// Enable logger
	app.Use(middlewares.Logger())

	// Enable Swagger
	app.Use(middlewares.Swagger())

	// Create route for "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber 👋!")
	})

	// Create route for "/api"
	api := app.Group("/api")

	// Create route for "/api/v1"
	v1 := api.Group("/v1")

	// Create routes for "/api/v1"
	routes.AuthRoute(v1)
	routes.UserRoute(v1)
	routes.PostRoute(v1)

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})

	// Start server
	var host = config.APP_HOST
	var port = config.APP_PORT

	err := app.Listen(host + ":" + port)

	if err != nil {
		return
	}
}
