package router

import (
	"absensibe/internal/student"
	"absensibe/middleware"

	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(api fiber.Router) {
	// Initialize dependencies
	repo := student.NewStudentRepository()
	service := student.NewStudentService(repo)
	handler := student.NewStudentHandler(service)

	// Student routes
	studentGroup := api.Group("/student")
	studentGroup.Use(middleware.AuthRequired())

	// Main endpoint
	studentGroup.Get("/info", handler.GetStudentInfo)
}
