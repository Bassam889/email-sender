package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leocm889/server/controllers"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/submit")
	v1.Post("/", controllers.SubmitFormHandler)
}
