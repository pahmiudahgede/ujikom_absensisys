package router

import (
	"absensibe/internal/attendance"
	"absensibe/middleware"

	"github.com/gofiber/fiber/v2"
)

func AttendanceRoutes(api fiber.Router) {

	repo := attendance.NewAttendanceRulesRepository()
	service := attendance.NewAttendanceRulesService(repo)
	handler := attendance.NewAttendanceRulesHandler(service)

	attendanceGroup := api.Group("/attendance")
	attendanceGroup.Use(middleware.AuthRequired())

	attendanceGroup.Get("/rules", handler.GetAttendanceRules)
}
