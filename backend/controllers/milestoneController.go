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

// ProjectCreate will add new clients to the database

func MilestoneCreate(c *fiber.Ctx) error {
	var body models.Milestone

	err := c.BodyParser(&body)

	if err != nil {
		return err
	}

	milestone := &models.Milestone{
		MilestoneName:    body.MilestoneName,
		MilestoneCode:    body.MilestoneCode,
		MilestoneStarted: body.MilestoneStarted,
		MilestonePrereq:  body.MilestonePrereq,
		MilestoneStart:   body.MilestoneStart,
		MilestoneEnd:     body.MilestoneEnd,
		MilestoneDetails: body.MilestoneDetails,
		ProjectID:        body.ProjectID,
	}

	database.DB.Create(&milestone)

	return c.JSON(milestone)
}
