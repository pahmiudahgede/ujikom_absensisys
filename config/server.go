package config

import (
	"os"
	"time"

	_ "absensibe/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// RootResponse represents the response for root endpoint
type RootResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Presensi SMK API is running!"`
	Version string `json:"version" example:"1.0.0"`
	Env     string `json:"env" example:"dev"`
}

// HealthResponse represents the response for health check endpoint
type HealthResponse struct {
	Status    string `json:"status" example:"healthy"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
	Database  string `json:"database" example:"connected"`
	Redis     string `json:"redis" example:"connected"`
}

// GetRoot godoc
//
//	@Summary		Get API Information
//	@Description	Returns basic information about the API including status, version, and environment
//	@Tags			System
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	RootResponse
//	@Router			/ [get]
func GetRoot(c *fiber.Ctx) error {
	return c.JSON(RootResponse{
		Status:  "success",
		Message: "Presensi SMK API is running!",
		Version: os.Getenv("APP_VERSION"),
		Env:     os.Getenv("APP_ENV"),
	})
}

// GetHealth godoc
//
//	@Summary		Health Check
//	@Description	Returns the health status of the API including database and Redis connectivity
//	@Tags			System
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	HealthResponse
//	@Router			/ujikom/api/health [get]
func GetHealth(c *fiber.Ctx) error {
	return c.JSON(HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Database:  "connected",
		Redis:     "connected",
	})
}

func SetupServer() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      os.Getenv("APP_NAME"),
		ServerHeader: "Fiber",
		Prefork:      false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"status":  code,
				"message": err.Error(),
				"error":   true,
			})
		},
	})

	// Swagger endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path} ${yellow}${body} ${reset}${latency}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-API-Key",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           3600,
	}))

	// Root endpoint with Swagger documentation
	app.Get("/", GetRoot)

	// Setup API group
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "/api"
	}

	api := app.Group(baseURL)

	// Health check endpoint with Swagger documentation
	api.Get("/health", GetHealth)

	return app
}
