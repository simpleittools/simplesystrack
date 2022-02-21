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

	// Project Routes
	project := app.Group("api/project")
	project.Get("/", controllers.ProjectIndex)
	project.Post("/", controllers.ProjectCreate)
	project.Get("/:project_code", controllers.ProjectShow)
	project.Patch("/:project_code", controllers.ProjectUpdate)
	project.Delete("/:project_code", controllers.ProjectDestroy)

}
