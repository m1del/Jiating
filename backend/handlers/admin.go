package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m1del/Jiating/backend/database"
	"github.com/m1del/Jiating/backend/models"
	"github.com/m1del/Jiating/backend/utils"
)

func CreateAdmin(c *fiber.Ctx) error {
	admin := new(models.Admin)
	if err := c.BodyParser(admin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot parse admin data",
		})
	}

	// hash password
	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot hash password",
		})
	}
	admin.PasswordHash = hashedPassword // store hashed password instead of plaintext

	// save admin to database
	database.DB.Db.Create(&admin)

	return c.Status(fiber.StatusCreated).JSON(admin)
}
