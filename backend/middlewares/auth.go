package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

/*
Middelwares are functions that server as prelude to actions
carried out by the API, e.g. validating that the user is logged in
or restricting access to certain routes to only admins.
*/

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}
