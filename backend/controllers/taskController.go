package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplesystrack/database"
	"github.com/simpleittools/simplesystrack/models"
)

// TaskIndex will collect all the active Tasks
func TaskIndex(c *fiber.Ctx) error {
	var tasks []models.Task

	database.DB.Joins("Project").Find(&tasks)

	return c.JSON(tasks)

}

// TaskCreate will add new tasks to the database
func TaskCreate(c *fiber.Ctx) error {
	var data models.Task

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	task := &models.Task{
		TaskName:    data.TaskName,
		TaskCode:    data.TaskCode,
		IsActive:    data.IsActive,
		TaskDetails: data.TaskDetails,
		MilestoneID: data.MilestoneID,
	}

	database.DB.Create(&task)

	return c.JSON(task)
}

// TaskShow will return the results of a selected Task and include the Milestone
func TaskShow(c *fiber.Ctx) error {
	taskCode := c.Params("task_code")
	task := models.Task{}
	database.DB.Joins("Milestone").Find(&task, "task_code", taskCode)

	return c.JSON(task)
}

// TaskUpdate will PATCH the Task details after edited by the user
func TaskUpdate(c *fiber.Ctx) error {
	taskCode := c.Params("task_code")

	var data models.Task

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	task := &models.Task{
		TaskName:    data.TaskName,
		TaskCode:    data.TaskCode,
		IsActive:    data.IsActive,
		TaskDetails: data.TaskDetails,
		MilestoneID: data.MilestoneID,
	}

	err = database.DB.Model(&models.Task{}).Where("task_code = ?", taskCode).Updates(&task).Error
	if err != nil {
		return err
	}

	return c.JSON(task)
}

// TaskDestroy will soft delete the task from the database
func TaskDestroy(c *fiber.Ctx) error {
	taskCode := c.Params("task_code")

	database.DB.Where("task_code = ?", taskCode).Delete(&models.Task{})

	return c.JSON(fiber.Map{
		"message": "Task Removed",
	})
}
