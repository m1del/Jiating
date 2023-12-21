package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/m1del/Jiating/backend/config"
	"github.com/m1del/Jiating/backend/database"
	"github.com/m1del/Jiating/backend/models"
	"github.com/m1del/Jiating/backend/utils"
)

// Login route
func Login(c *fiber.Ctx) error {
	// extract the credentials from the request body
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// check if the user exists
	admin := new(models.Admin)
	database.DB.Db.Where("email = ?", loginRequest.Email).First(&admin)
	if admin.ID == 0 || !utils.CheckPasswordHash(loginRequest.Password, admin.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	claims := jtoken.MapClaims{
		"sub":   admin.AdminID,                         // Unique identifier for the user
		"email": admin.Email,                           // Email address of the admin
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":   time.Now().Unix(),                     // Issued at current time
		"iss":   "jiating",                             // Issuer
		"aud":   "jiating-admins",                      // Audience
	}

	// create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

// protected route
func Protected(c *fiber.Ctx) error {
	// get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	favPhrase := claims["fav"].(string)
	return c.SendString("Welcome 👋" + email + " " + favPhrase)
}
