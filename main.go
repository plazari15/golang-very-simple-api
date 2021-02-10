package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/palazari19/gotestefiber/routes"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	routes.Register(api)
}
