package main

import (
	"github.com/JulianElisii/Crud-Go-/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.UseMoviesRoute(app)

	app.Listen(":4000")
}
