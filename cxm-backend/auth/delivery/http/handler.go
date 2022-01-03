package http

import (
	"cxm-auth/auth"
	"cxm-auth/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

type CreateUserDTO struct {
	Username string `form:"username" validate:"required,min=6,max=32"`
	Password string `form:"password" validate:"required,min=6,max=32"`
}

func (h *Handler) SignUp(c *fiber.Ctx) error {
	data := new(CreateUserDTO)

	if err := c.BodyParser(data); err != nil {
		return helper.SimpleError(c, err)
	}

	if err := validator.New().Struct(data); err != nil {
		return helper.SimpleError(c, err)
	}

	if err := h.useCase.SignUp(c.UserContext(), data.Username, data.Password); err != nil {
		return helper.SimpleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"msg": "OK",
	})
}

type SignInDTO struct {
	Username string `form:"username" validate:"required,min=6,max=32"`
	Password string `form:"password" validate:"required,min=6,max=32"`
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	data := new(SignInDTO)

	if err := c.BodyParser(data); err != nil {
		return helper.SimpleError(c, err)
	}

	if err := validator.New().Struct(data); err != nil {
		return helper.SimpleError(c, err)
	}

	token, err := h.useCase.SignIn(c.UserContext(), data.Username, data.Password)

	if err != nil {
		return helper.SimpleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"token": token,
	})
}
