package student

import (
	"absensibe/middleware"
	"absensibe/utils"
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler interface {
	GetStudentInfo(c *fiber.Ctx) error
}

type studentHandler struct {
	service StudentService
}

func NewStudentHandler(service StudentService) StudentHandler {
	return &studentHandler{
		service: service,
	}
}

// GetStudentInfo godoc
//
//	@Summary		Get Student Information
//	@Description	Get student information with optional relations (class, attendance, schedule)
//	@Tags			Student
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Security		BearerAuth
//	@Param			class		query		bool	false	"Include class information"
//	@Param			absen		query		bool	false	"Include attendance information"
//	@Param			schedule	query		bool	false	"Include schedule information"
//	@Param			month		query		string	false	"Filter attendance by month (format: 2024-01)"
//	@Param			limit		query		int		false	"Attendance records limit (default: 10, max: 50)"
//	@Success		200			{object}	utils.SuccessAPIResponse{data=StudentInfoResponse}
//	@Failure		400			{object}	utils.BadRequestAPIResponse
//	@Failure		401			{object}	utils.UnauthorizedAPIResponse
//	@Failure		404			{object}	utils.NotFoundAPIResponse
//	@Failure		500			{object}	utils.InternalServerErrorAPIResponse
//	@Router			/ujikom/api/student/info [get]
func (h *studentHandler) GetStudentInfo(c *fiber.Ctx) error {
	// Get current user from middleware
	user := middleware.GetUserFromContext(c)
	if user == nil {
		return utils.Unauthorized(c)
	}

	// Parse request
	req := &StudentInfoRequest{
		IncludeClass:      c.Query("class") != "",
		IncludeAttendance: c.Query("absen") != "",
		IncludeSchedule:   c.Query("schedule") != "",
		AttendanceMonth:   c.Query("month"),
	}

	// Parse limit
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			if limit > 50 {
				limit = 50 // Max limit
			}
			req.AttendanceLimit = limit
		}
	}

	// Validate month format if provided
	if req.AttendanceMonth != "" {
		if _, err := time.Parse("2006-01", req.AttendanceMonth); err != nil {
			return utils.BadRequest(c, "Invalid month format. Use YYYY-MM (e.g., 2024-01)")
		}
	}

	// Auto-include class if schedule is requested
	if req.IncludeSchedule {
		req.IncludeClass = true
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get student info
	studentInfo, err := h.service.GetStudentInfo(ctx, user.UserID, req)
	if err != nil {
		if err.Error() == "student not found" {
			return utils.NotFound(c, "Student")
		}
		return utils.InternalError(c)
	}

	return utils.OK(c, studentInfo)
}
