package services

import (
	"go-rest-api/src/config"
	"go-rest-api/src/models"

	"github.com/gofiber/fiber/v3"
)

func GetAllUsers(c fiber.Ctx) error {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	return c.JSON(users)
}

func GetUser(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.JSON(user)
}

func CreateUser(c fiber.Ctx) error {
	user := new(models.User)
	if err := c.Bind().JSON(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	var updateData models.User
	if err := c.Bind().JSON(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	config.DB.Model(&user).Updates(updateData)
	return c.JSON(user)
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	config.DB.Delete(&user)
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
