package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplesystrack/database"
	"github.com/simpleittools/simplesystrack/models"
)

// ClientIndex will show all clients in the database
func ClientIndex(c *fiber.Ctx) error {
	var clients []models.Client

	database.DB.Where("is_active = ?", "true").Find(&clients)
	//database.DB.Preload(clause.Associations).Where("is_active = 1").Find(&clients)

	return c.JSON(clients)
}

// ClientCreate will add new clients to the database
func ClientCreate(c *fiber.Ctx) error {
	var body models.Client

	err := c.BodyParser(&body)
	// ToDo Error handling is not returning errors. Need to look into GORM error handling
	if err != nil {
		return err
	}

	client := &models.Client{
		ClientName:     body.ClientName,
		ClientInitials: body.ClientInitials,
		ClientStreet:   body.ClientStreet,
		ClientStreet2:  body.ClientStreet2,
		ClientZip:      body.ClientZip,
		ClientState:    body.ClientState,
		IsActive:       body.IsActive,
		Comments:       body.Comments,
	}

	database.DB.Create(&client)

	return c.JSON(client)

}

// ClientShow will return the results of a selected client
func ClientShow(c *fiber.Ctx) error {
	clientInitials := c.Params("client_initials")
	client := models.Client{}
	database.DB.Preload("Projects", "project_is_active=?", true).Find(&client, "client_initials", clientInitials)

	return c.JSON(client)
}

// ClientUpdate will PATCH the client details after edited by the user
func ClientUpdate(c *fiber.Ctx) error {
	clientInitials := c.Params("client_initials")

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	client := models.Client{
		ClientName:     data["client_name"],
		ClientInitials: data["client_initials"],
	}

	database.DB.Model(&models.Client{}).Where("client_initials = ?", clientInitials).Updates(map[string]interface{}{
		"client_name":     client.ClientName,
		"client_initials": client.ClientInitials,
	})

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// ClientDestroy will remove the client from the database
func ClientDestroy(c *fiber.Ctx) error {
	clientInitials := c.Params("client_initials")

	database.DB.Where("client_initials = ?", clientInitials).Delete(&models.Client{})

	return c.JSON(fiber.Map{
		"message": "Client Removed",
	})
}
