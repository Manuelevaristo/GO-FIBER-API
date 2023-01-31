package main

import (
	"api/configs"
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//executar banco de dados
	configs.ConnectDB()

	//rotas
	routes.UserRoute(app)

	app.Listen(":8080")
}
