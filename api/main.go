package main

import (
	"api/configs"
	"api/routes" //add this

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app) //add this

	app.Listen(":8080")
}