package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplesystrack/database"
	"github.com/simpleittools/simplesystrack/models"
)

// ProjectIndex will show all projects in the database
func ProjectIndex(c *fiber.Ctx) error {
	var projects []models.Project

	database.DB.Joins("Client").Preload("Milestones").Find(&projects)

	return c.JSON(projects)
}

// ProjectCreate will add new clients to the database

func ProjectCreate(c *fiber.Ctx) error {
	var body models.Project

	err := c.BodyParser(&body)

	if err != nil {
		return err
	}

	project := &models.Project{
		ProjectName:        body.ProjectName,
		ProjectCode:        body.ProjectCode,
		ClientID:           body.ClientID,
		ProjectDescription: body.ProjectDescription,
		ProjectBudgetHrs:   body.ProjectBudgetHrs,
		ProjectBudget:      body.ProjectBudget,
		ProjectIsActive:    body.ProjectIsActive,
		ProjectComments:    body.ProjectComments,
	}

	database.DB.Create(&project)

	return c.JSON(project)
}

// ProjectShow will return the results of a selected client
func ProjectShow(c *fiber.Ctx) error {
	ProjectCode := c.Params("project_code")
	project := models.Project{}
	database.DB.Preload("Milestones", "milestone_is_active=?", true).Joins("Client").Find(&project, "project_code", ProjectCode)

	return c.JSON(project)
}

// ProjectUpdate will PATCH the client details after edited by the user
func ProjectUpdate(c *fiber.Ctx) error {
	ProjectCode := c.Params("project_code")

	var data models.Project

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	project := &models.Project{
		ProjectName: data.ProjectName,
		ProjectCode: data.ProjectCode,
		ClientID:    data.ClientID,
	}

	database.DB.Model(&models.Project{}).Where("project_code = ?", ProjectCode).Updates(map[string]interface{}{
		"project_name": project.ProjectName,
		"project_code": project.ProjectCode,
		"client_id":    project.ClientID,
	})

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// ProjectDestroy will remove the client from the database
func ProjectDestroy(c *fiber.Ctx) error {
	ProjectCode := c.Params("project_code")

	database.DB.Where("project_code = ?", ProjectCode).Delete(&models.Project{})

	return c.JSON(fiber.Map{
		"message": "Project Removed",
	})
}
