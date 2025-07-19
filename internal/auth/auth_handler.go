package auth

import (
	"absensibe/middleware"
	"absensibe/utils"
	"context"
	"strings"
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

// Login godoc
//
//	@Summary		Student Login
//	@Description	Authenticate student dengan NISN/NIS dan password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest							true	"Login credentials"
//	@Success		200		{object}	utils.LoginSuccessResponse				"Login successful"
//	@Failure		400		{object}	utils.BadRequestAPIResponse				"Bad request - invalid input"
//	@Failure		401		{object}	utils.UnauthorizedAPIResponse			"Unauthorized - invalid credentials"
//	@Failure		500		{object}	utils.InternalServerErrorAPIResponse	"Internal server error"
//	@Security		ApiKeyAuth
//	@Router			/student/login [post]
func (h *StudentHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	// Additional validation
	if strings.TrimSpace(req.Identifier) == "" {
		return utils.BadRequest(c, "NISN/NIS is required")
	}
	
	if strings.TrimSpace(req.Password) == "" {
		return utils.BadRequest(c, "Password is required")
	}

	// Auto-fill device info if not provided
	if req.DeviceInfo == "" {
		req.DeviceInfo = c.Get("User-Agent")
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

// Logout godoc
//
//	@Summary		Student Logout
//	@Description	Logout student dan hapus session yang aktif
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.SuccessOnlyAPIResponse			"Logout successful"
//	@Failure		401	{object}	utils.UnauthorizedAPIResponse			"Unauthorized - invalid or expired token"
//	@Failure		500	{object}	utils.InternalServerErrorAPIResponse	"Internal server error"
//	@Security		ApiKeyAuth
//	@Security		BearerAuth
//	@Router			/student/logout [post]
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