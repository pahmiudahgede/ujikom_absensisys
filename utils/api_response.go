package utils

import (
	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	CurrentPage int   `json:"current_page,omitempty"`
	PerPage     int   `json:"per_page,omitempty"`
	Total       int64 `json:"total,omitempty"`
	LastPage    int   `json:"last_page,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, err interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func BadRequestResponse(c *fiber.Ctx, message string, err interface{}) error {
	return ErrorResponse(c, fiber.StatusBadRequest, message, err)
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
	return c.Status(fiber.StatusUnprocessableEntity).JSON(APIResponse{
		Success: false,
		Message: message,
		Error:   errors,
	})
}

func InternalServerErrorResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusInternalServerError, message, nil)
}

func PaginatedResponse(c *fiber.Ctx, message string, data interface{}, meta *Meta) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func NoContentResponse(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func CreatePaginationMeta(currentPage, perPage int, total int64) *Meta {
	lastPage := int((total + int64(perPage) - 1) / int64(perPage))
	if lastPage < 1 {
		lastPage = 1
	}

	return &Meta{
		CurrentPage: currentPage,
		PerPage:     perPage,
		Total:       total,
		LastPage:    lastPage,
	}
}

func CreateValidationError(field, message, code string) ErrorDetail {
	return ErrorDetail{
		Field:   field,
		Message: message,
		Code:    code,
	}
}

func OK(c *fiber.Ctx, data interface{}) error {
	return SuccessResponse(c, "Success", data)
}

func Created(c *fiber.Ctx, data interface{}) error {
	return CreatedResponse(c, "Data created successfully", data)
}

func Updated(c *fiber.Ctx, data interface{}) error {
	return SuccessResponse(c, "Data updated successfully", data)
}

func Deleted(c *fiber.Ctx) error {
	return SuccessResponse(c, "Data deleted successfully", nil)
}

func BadRequest(c *fiber.Ctx, message string) error {
	return BadRequestResponse(c, message, nil)
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
