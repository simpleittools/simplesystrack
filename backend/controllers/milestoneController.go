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

// MilestoneShow will return the results of a selected Milestone and include the Project
func MilestoneShow(c *fiber.Ctx) error {
	milestoneCode := c.Params("milestone_code")
	milestone := models.Milestone{}
	database.DB.Joins("Project").Find(&milestone, "milestone_code", milestoneCode)

	return c.JSON(milestone)
}

// MilestoneUpdate will PATCH the Milestone details after edited by the user
func MilestoneUpdate(c *fiber.Ctx) error {
	MilestoneCode := c.Params("milestone_code")

	var data models.Milestone

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	milestone := &models.Milestone{
		MilestoneName:    data.MilestoneName,
		MilestoneCode:    data.MilestoneCode,
		MilestoneStarted: data.MilestoneStarted,
		MilestonePrereq:  data.MilestonePrereq,
		MilestoneStart:   data.MilestoneStart,
		MilestoneEnd:     data.MilestoneEnd,
		MilestoneDetails: data.MilestoneDetails,
		ProjectID:        data.ProjectID,
	}

	database.DB.Model(&models.Milestone{}).Where("milestone_code = ?", MilestoneCode).Updates(map[string]interface{}{
		"milestone_name":    milestone.MilestoneName,
		"milestone_code":    milestone.MilestoneCode,
		"milestone_started": milestone.MilestoneStarted,
		"milestone_prereq":  milestone.MilestonePrereq,
		"milestone_start":   milestone.MilestoneStart,
		"milestone_end":     milestone.MilestoneEnd,
		"milestone_details": milestone.MilestoneDetails,
		"project_id":        milestone.ProjectID,
	})

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// MilestoneDestroy will remove the milestone from the database
func MilestoneDestroy(c *fiber.Ctx) error {
	MilestoneCode := c.Params("milestone_code")

	database.DB.Where("milestone_code = ?", MilestoneCode).Delete(&models.Milestone{})

	return c.JSON(fiber.Map{
		"message": "Milestone Removed",
	})
}
