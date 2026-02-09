package routes

import (
	"go-rest-api/src/services"

	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	users := api.Group("/users")

	users.Get("/", services.GetAllUsers)
	users.Get("/:id", services.GetUser)
	users.Post("/", services.CreateUser)
	users.Put("/:id", services.UpdateUser)
	users.Delete("/:id", services.DeleteUser)
}
