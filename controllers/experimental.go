package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func ExperimentalIndex(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "Hello World from Controller :D",
	})
}
