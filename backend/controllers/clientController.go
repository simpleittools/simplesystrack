package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplesystrack/database"
	"github.com/simpleittools/simplesystrack/models"
)

func CreateClient(c *fiber.Ctx) {
	db := database.Conn

	client := new(models.Client)
	err := c.BodyParser(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Create(&client)

	c.JSON(client)

}
