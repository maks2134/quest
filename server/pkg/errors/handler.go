package errors

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

func HandlerErrorFormatter(ctx fiber.Ctx, err error) error {
	var e *Error
	var fe *fiber.Error
	if errors.As(err, &e) {
		return ctx.Status(e.StatusCode).JSON(e)
	} else if errors.As(err, &fe) {
		errorResponse := NewSimpleError(fe.Code, fe.Error())
		return ctx.Status(fe.Code).JSON(errorResponse)
	}
	errorResponse := NewSimpleError(fiber.StatusInternalServerError, "Internal server error")
	return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse)
}
