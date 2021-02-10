package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/palazari19/gotestefiber/controllers"
)

func Register(app fiber.Router) {
	app.Get("/", controllers.ExperimentalIndex)
}

func index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "You are at the endpoint 😉",
	})
}
