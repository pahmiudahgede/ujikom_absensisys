package router

import (
	"absensibe/config"
	"absensibe/internal/auth"
	"absensibe/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(api fiber.Router) {
	studentRepo := auth.NewStudentRepository(config.DB)
	studentService := auth.NewStudentService(studentRepo)
	studentHandler := auth.NewStudentHandler(studentService)

	studentAuth := api.Group("/student")
	studentAuth.Post("/login", studentHandler.Login)
	studentAuth.Post("/logout", middleware.AuthRequired(), studentHandler.Logout)
}
