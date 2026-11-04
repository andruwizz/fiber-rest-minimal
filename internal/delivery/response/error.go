package response

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"

	"github.com/andruwizz/fiber-rest-minimal/internal/entity"
	"github.com/go-playground/validator/v10"
)

type errorDetail struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   errorDetail `json:"error"`
}

func extractValidationError(ve validator.ValidationErrors) (string, map[string]string) {
	errs := map[string]string{}
	for _, f := range ve {
		err := strings.Split(f.Error(), "Error:")
		errs[strings.ToLower(f.Field())] = fmt.Sprintf("%s.", err[1])
	}

	message := fmt.Sprintf("%s.", strings.Split(ve[0].Error(), "Error:")[1])
	if errCount := len(ve); errCount > 1 {
		message = fmt.Sprintf("%s (and %d more)", message, errCount-1)
	}

	return message, errs
}

func ErrorValidation(ctx fiber.Ctx, code int, err error) error {
	if ierr, ok := err.(*fiber.BindError); ok {
		err = ierr.Err
		res := ErrorResponse{
			Success: false,
			Error: errorDetail{
				Message: "request body is invalid",
			},
		}
		return ctx.Status(code).JSON(res)
	}

	msg, errs := extractValidationError(err.(validator.ValidationErrors))
	res := ErrorResponse{
		Success: false,
		Error: errorDetail{
			Message: msg,
			Errors:  errs,
		},
	}

	return ctx.Status(code).JSON(res)
}

func errorResource(ctx fiber.Ctx, code int, err error) error {
	res := ErrorResponse{
		Success: false,
		Error: errorDetail{
			Message: err.Error(),
		},
	}

	return ctx.Status(code).JSON(res)
}

func ErrorResource(c fiber.Ctx, err error) error {
	if errors.Is(err, entity.ErrRecordNotFound) {
		return errorResource(c, fiber.StatusNotFound, err)
	}
	return errorResource(c, fiber.StatusInternalServerError, err)
}
