package authhttp

import (
	"cxm-auth/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CurrentUser(uc auth.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		currentUser, err := uc.GetUser(c.UserContext(), username)
		if err != nil {
			return err
		}

		c.Locals("currentUser", currentUser)

		return c.Next()
	}
}
