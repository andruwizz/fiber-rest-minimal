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
		parts := strings.Split(f.Error(), "Error:")
		errText := f.Error()
		if len(parts) > 1 {
			errText = strings.TrimSpace(parts[1])
		}
		errs[strings.ToLower(f.Field())] = fmt.Sprintf("%s.", strings.TrimSuffix(errText, "."))
	}

	firstPart := ve[0].Error()
	parts := strings.Split(firstPart, "Error:")
	if len(parts) > 1 {
		firstPart = strings.TrimSpace(parts[1])
	}
	message := fmt.Sprintf("%s.", strings.TrimSuffix(firstPart, "."))
	if errCount := len(ve); errCount > 1 {
		message = fmt.Sprintf("%s (and %d more)", message, errCount-1)
	}

	return message, errs
}

func ErrorValidation(ctx fiber.Ctx, code int, err error) error {
	var bindErr *fiber.BindError
	if errors.As(err, &bindErr) {
		err = bindErr.Err
	}

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		msg, errs := extractValidationError(ve)
		res := ErrorResponse{
			Success: false,
			Error: errorDetail{
				Message: msg,
				Errors:  errs,
			},
		}
		return ctx.Status(code).JSON(res)
	}

	res := ErrorResponse{
		Success: false,
		Error: errorDetail{
			Message: "request payload is invalid",
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
