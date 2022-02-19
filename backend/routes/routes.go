package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplesystrack/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Index)

	app.Post("/api/client", controllers.CreateClient)
}
