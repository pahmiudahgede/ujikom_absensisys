package auth

import (
	"absensibe/middleware"
	"absensibe/utils"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	service   StudentService
	validator *validator.Validate
}

func NewStudentHandler(service StudentService) *StudentHandler {
	return &StudentHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *StudentHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return utils.BadRequest(c, "Validation failed")
	}

	loginResponse, err := h.service.Login(ctx, req, c)
	if err != nil {
		return utils.UnauthorizedResponse(c, err.Error())
	}

	return utils.SuccessResponse(c, "login success", loginResponse)
}

func (h *StudentHandler) Logout(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := middleware.GetUserFromContext(c)
	if user == nil {
		return utils.Unauthorized(c)
	}

	err := h.service.Logout(ctx, user.SessionID)
	if err != nil {
		return utils.InternalError(c)
	}

	return utils.SuccessResponseOnly(c, "logout success")
}
