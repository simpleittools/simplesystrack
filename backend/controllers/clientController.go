package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplesystrack/database"
	"github.com/simpleittools/simplesystrack/models"
)

func CreateClient(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	client := models.Client{
		ClientName:     data["client_name"],
		ClientInitials: data["client_initials"],
	}

	database.DB.Create(&client)

	return c.JSON(client)

}
