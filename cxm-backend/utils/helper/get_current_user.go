package helper

import (
	"cxm-auth/module/auth"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx) *auth.User {
	currentUser := c.Locals("currentUser").(*auth.User)

	return currentUser
}
