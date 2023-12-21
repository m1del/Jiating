package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m1del/Jiating/backend/database"
	"github.com/m1del/Jiating/backend/models"
	"github.com/m1del/Jiating/backend/utils"
)

// temp struct to parse admin creation request body
type CreateAdminRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Position string `json:"position"`
	Status   string `json:"status"`
}

func CreateAdmin(c *fiber.Ctx) error {
	var request CreateAdminRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot parse admin data",
		})
	}

	// hash password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot hash password",
		})
	}

	// create admin
	admin := models.Admin{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: hashedPassword,
		Position:     request.Position,
		Status:       request.Status,
	}

	// save admin to database
	database.DB.Db.Create(&admin)

	return c.Status(fiber.StatusCreated).JSON(admin)
}

func AuthenticateAdmin(c *fiber.Ctx) error {
	// struct to parse login request body
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	request := new(LoginReq)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot parse login request body",
		})
	}
	admin := new(models.Admin)

	// find admin by email
	database.DB.Db.Where("email = ?", request.Email).First(&admin)

	if admin.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Admin not found",
		})
	}

	if !utils.CheckPasswordHash(request.Password, admin.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	// generate jwt token

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}
