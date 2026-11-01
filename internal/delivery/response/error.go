package response

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"

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

func ErrorValidation(ctx fiber.Ctx, code int, err error) {
	msg, errs := extractValidationError(err.(validator.ValidationErrors))
	res := ErrorResponse{
		Success: false,
		Error: errorDetail{
			Message: msg,
			Errors:  errs,
		},
	}

	ctx.Status(code).JSON(res)
}

func ErrorResource(ctx fiber.Ctx, code int, err error) {
	res := ErrorResponse{
		Success: false,
		Error: errorDetail{
			Message: err.Error(),
		},
	}

	ctx.Status(code).JSON(res)
}
