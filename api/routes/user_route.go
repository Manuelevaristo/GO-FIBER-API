package routes

import (
	"api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/api/v1/user", controllers.CreateUser)
	app.Get("/api/v1/user/:userId", controllers.GetAUser)
	app.Put("/api/v1/user/:userId", controllers.EditAUser)
	app.Delete("/api/v1/user/:userId", controllers.DeleteAUser)
	app.Get("/api/v1/users", controllers.GetAllUsers)
}
