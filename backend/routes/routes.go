package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplesystrack/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Index)

	// Client Routes
	client := app.Group("api/client")
	client.Get("/", controllers.ClientIndex)
	//client.Get("/new", controllers.ClientNew)
	client.Post("/", controllers.ClientCreate)
	client.Get("/:client_initials", controllers.ClientShow)
	//client.Get("/:clientInitials/edit", controllers.ClientEdit)
	client.Patch("/:client_initials", controllers.ClientUpdate)
	client.Delete("/:client_initials", controllers.ClientDestroy)
	// ToDo: frontend needs to also have ClientEdit and ClientNew

	// Project Routes
	project := app.Group("api/project")
	project.Get("/", controllers.ProjectIndex)
	project.Post("/", controllers.ProjectCreate)
	project.Get("/:project_code", controllers.ProjectShow)
	project.Patch("/:project_code", controllers.ProjectUpdate)
	project.Delete("/:project_code", controllers.ProjectDestroy)
	// ToDo: frontend needs to also have ProjectEdit and ProjectNew

	/*
		ToDo:
			Milestones -> need Model, Controller, Milestones belong to Projects but should return Tasks,
			Project, PM, % complete (closedTasks/totalTasks), hours used, % to budget (hours used/total budget), Client
	*/
	// Milestones Routes
	milestone := app.Group("api/milestone")
	milestone.Get("/", controllers.MilestoneIndex)
	milestone.Post("/", controllers.MilestoneCreate)

	/*
		ToDo:
			Tasks -> need Model, Controller,
			Tasks belong to milestones but should return Milestone, Project, PM, Assigned to (user), Client
	*/

	/*
		ToDo:
			Users -> Model, controller, authentication, many->many roles, hasMany Tasks
	*/

	/*
		ToDo:
			Roles -> Model, controller, many->many users
	*/

	/*
		Todo:
			Permissions -> Model, controller, many->many roles
			FrontEnd -> Will needs a permissions viewer
	*/

}
