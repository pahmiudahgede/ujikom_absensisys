package attendance

import (
	"absensibe/middleware"
	"absensibe/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AttendanceRulesHandler interface {
	GetAttendanceRules(c *fiber.Ctx) error
}

type attendanceRulesHandler struct {
	service AttendanceRulesService
}

func NewAttendanceRulesHandler(service AttendanceRulesService) AttendanceRulesHandler {
	return &attendanceRulesHandler{
		service: service,
	}
}

// GetAttendanceRules godoc
//
//	@Summary		Get Attendance Rules and Safe Areas
//	@Description	Get attendance settings, safe areas, and current attendance status for student
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Security		BearerAuth
//	@Success		200	{object}	utils.SuccessAPIResponse{data=AttendanceRulesResponse}
//	@Failure		401	{object}	utils.UnauthorizedAPIResponse
//	@Failure		404	{object}	utils.NotFoundAPIResponse
//	@Failure		500	{object}	utils.InternalServerErrorAPIResponse
//	@Router			/ujikom/api/attendance/rules [get]
func (h *attendanceRulesHandler) GetAttendanceRules(c *fiber.Ctx) error {
	// Get current user from middleware
	user := middleware.GetUserFromContext(c)
	if user == nil {
		return utils.Unauthorized(c)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get attendance rules
	rules, err := h.service.GetAttendanceRules(ctx, user.UserID)
	if err != nil {
		if err.Error() == "student not found" {
			return utils.NotFound(c, "Student")
		}
		if err.Error() == "attendance settings not found for school" {
			return utils.NotFound(c, "Attendance settings")
		}
		return utils.InternalError(c)
	}

	return utils.OK(c, rules)
}
