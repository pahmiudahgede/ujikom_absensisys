package utils

import (
	"github.com/gofiber/fiber/v2"
)

// APIResponse represents the standard API response structure
type APIResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data,omitempty"`
} //	@name	APIResponse

// MetaResponse represents metadata for API responses
type MetaResponse struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"Success"`
	Page    *int   `json:"page,omitempty" example:"1"`
	Limit   *int   `json:"limit,omitempty" example:"10"`
	Total   *int64 `json:"total,omitempty" example:"100"`
} //	@name	MetaResponse

// ErrorDetail represents validation error detail
type ErrorDetail struct {
	Field   string `json:"field,omitempty" example:"email"`
	Message string `json:"message" example:"Email is required"`
	Code    string `json:"code,omitempty" example:"required"`
} //	@name	ErrorDetail

// LoginSuccessResponse represents successful login response
type LoginSuccessResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"200"`
		Message string `json:"message" example:"login success"`
	} `json:"meta"`
	Data struct {
		UserID       string `json:"user_id" example:"003061ee-97ff-4d00-8155-f4bf15e319dd"`
		NISN         string `json:"nisn" example:"2024000579"`
		Name         string `json:"name" example:"Eko Saputra"`
		SessionID    string `json:"session_id" example:"b5e882d0eb21c5ba7752d2b8f216ccad"`
		AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
		CreatedAt    string `json:"created_at" example:"0001-01-01T00:00:00Z"`
		LastActivity string `json:"last_activity" example:"0001-01-01T00:00:00Z"`
		DeviceInfo   string `json:"device_info" example:"postmanstudent123x0=="`
	} `json:"data"`
} //	@name	LoginSuccessResponse

// SuccessAPIResponse represents successful response with data
type SuccessAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"200"`
		Message string `json:"message" example:"Success"`
	} `json:"meta"`
	Data interface{} `json:"data"`
} //	@name	SuccessAPIResponse

// SuccessOnlyAPIResponse represents successful response without data
type SuccessOnlyAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"200"`
		Message string `json:"message" example:"Operation successful"`
	} `json:"meta"`
} //	@name	SuccessOnlyAPIResponse

// CreatedAPIResponse represents created response
type CreatedAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"201"`
		Message string `json:"message" example:"Data created successfully"`
	} `json:"meta"`
	Data interface{} `json:"data"`
} //	@name	CreatedAPIResponse

// PaginatedAPIResponse represents paginated response
type PaginatedAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"200"`
		Message string `json:"message" example:"Data retrieved successfully"`
		Page    int    `json:"page" example:"1"`
		Limit   int    `json:"limit" example:"10"`
		Total   int64  `json:"total" example:"100"`
	} `json:"meta"`
	Data interface{} `json:"data"`
} //	@name	PaginatedAPIResponse

// BadRequestAPIResponse represents 400 bad request response
type BadRequestAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"400"`
		Message string `json:"message" example:"Bad request"`
	} `json:"meta"`
	Data interface{} `json:"data,omitempty"`
} //	@name	BadRequestAPIResponse

// UnauthorizedAPIResponse represents 401 unauthorized response
type UnauthorizedAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"401"`
		Message string `json:"message" example:"Unauthorized access"`
	} `json:"meta"`
} //	@name	UnauthorizedAPIResponse

// ForbiddenAPIResponse represents 403 forbidden response
type ForbiddenAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"403"`
		Message string `json:"message" example:"Access forbidden"`
	} `json:"meta"`
} //	@name	ForbiddenAPIResponse

// NotFoundAPIResponse represents 404 not found response
type NotFoundAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"404"`
		Message string `json:"message" example:"Resource not found"`
	} `json:"meta"`
} //	@name	NotFoundAPIResponse

// ValidationErrorAPIResponse represents 422 validation error response
type ValidationErrorAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"422"`
		Message string `json:"message" example:"Validation failed"`
	} `json:"meta"`
	Data []ErrorDetail `json:"data"`
} //	@name	ValidationErrorAPIResponse

// InternalServerErrorAPIResponse represents 500 internal server error response
type InternalServerErrorAPIResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"500"`
		Message string `json:"message" example:"Internal server error"`
	} `json:"meta"`
} //	@name	InternalServerErrorAPIResponse

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
