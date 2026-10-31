package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func NewFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	})
	return app
}
