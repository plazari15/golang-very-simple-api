package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/palazari19/gotestefiber/models"
	"github.com/palazari19/gotestefiber/tools"
	"go.mongodb.org/mongo-driver/bson"
)

func ExperimentalIndex(ctx *fiber.Ctx) error {
	cursor, err := tools.CreateNewClient("flags").Db.Collection("squads").Find(ctx.Context(), bson.D{})

	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var features []models.Squad = make([]models.Squad, 0)

	//Quero entender mais desse cursor
	if err := cursor.All(ctx.Context(), &features); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.JSON(features)

	return ctx.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "Hello World from Controller :D",
	})
}
