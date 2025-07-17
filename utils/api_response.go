package utils

import (
	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data,omitempty"`
}

type MetaResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Page    *int   `json:"page,omitempty"`
	Limit   *int   `json:"limit,omitempty"`
	Total   *int64 `json:"total,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Meta: MetaResponse{
			Status:  200,
			Message: message,
		},
		Data: data,
	})
}

func SuccessResponseOnly(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Meta: MetaResponse{
			Status:  200,
			Message: message,
		},
	})
}

func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(APIResponse{
		Meta: MetaResponse{
			Status:  201,
			Message: message,
		},
		Data: data,
	})
}

func PaginatedResponse(c *fiber.Ctx, message string, data interface{}, page, limit int, total int64) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Meta: MetaResponse{
			Status:  200,
			Message: message,
			Page:    &page,
			Limit:   &limit,
			Total:   &total,
		},
		Data: data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := APIResponse{
		Meta: MetaResponse{
			Status:  statusCode,
			Message: message,
		},
	}

	if data != nil {
		response.Data = data
	}

	return c.Status(statusCode).JSON(response)
}

func BadRequestResponse(c *fiber.Ctx, message string, errors ...interface{}) error {
	var errorData interface{}
	if len(errors) > 0 && errors[0] != nil {
		errorData = errors[0]
	}
	return ErrorResponse(c, fiber.StatusBadRequest, message, errorData)
}

func UnauthorizedResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusUnauthorized, message, nil)
}

func ForbiddenResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusForbidden, message, nil)
}

func NotFoundResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusNotFound, message, nil)
}

func ValidationErrorResponse(c *fiber.Ctx, message string, errors []ErrorDetail) error {
	return ErrorResponse(c, fiber.StatusUnprocessableEntity, message, errors)
}

func InternalServerErrorResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusInternalServerError, message, nil)
}

func OK(c *fiber.Ctx, data interface{}) error {
	return SuccessResponse(c, "Success", data)
}

func OKOnly(c *fiber.Ctx, message string) error {
	return SuccessResponseOnly(c, message)
}

func Created(c *fiber.Ctx, data interface{}) error {
	return CreatedResponse(c, "Data created successfully", data)
}

func Updated(c *fiber.Ctx, data interface{}) error {
	return SuccessResponse(c, "Data updated successfully", data)
}

func Deleted(c *fiber.Ctx) error {
	return SuccessResponseOnly(c, "Data deleted successfully")
}

func BadRequest(c *fiber.Ctx, message string) error {
	return BadRequestResponse(c, message)
}

func Unauthorized(c *fiber.Ctx) error {
	return UnauthorizedResponse(c, "Unauthorized access")
}

func Forbidden(c *fiber.Ctx) error {
	return ForbiddenResponse(c, "Access forbidden")
}

func NotFound(c *fiber.Ctx, resource string) error {
	return NotFoundResponse(c, resource+" not found")
}

func InternalError(c *fiber.Ctx) error {
	return InternalServerErrorResponse(c, "Internal server error")
}

func CreateValidationError(field, message, code string) ErrorDetail {
	return ErrorDetail{
		Field:   field,
		Message: message,
		Code:    code,
	}
}

func NoContentResponse(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
