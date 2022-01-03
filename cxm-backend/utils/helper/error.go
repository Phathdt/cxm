package helper

import (
	"cxm-auth/utils/validator"

	"github.com/gofiber/fiber/v2"
)

func SimpleError(c *fiber.Ctx, err error) error {
	resp := validator.ToErrResponse(err)

	if resp == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(resp)
}
