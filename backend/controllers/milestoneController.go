package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplesystrack/database"
	"github.com/simpleittools/simplesystrack/models"
)

// MilestoneIndex will collect all the active milestones
func MilestoneIndex(c *fiber.Ctx) error {
	var milestones []models.Milestone

	database.DB.Joins("Project").Find(&milestones)

	return c.JSON(milestones)

}
