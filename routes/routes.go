package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/palazari19/gotestefiber/controllers"
)

func Register(app fiber.Router) {
	app.Get("/", controllers.ExperimentalIndex)
}
