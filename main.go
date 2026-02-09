package main

import (
	"go-rest-api/src/config"
	"go-rest-api/src/models"
	"go-rest-api/src/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	app := fiber.New()

	routes.SetupUserRoutes(app)

	app.Listen(":8081")
}
