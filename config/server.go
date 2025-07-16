package config

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Presensi SMK API is running!",
			"version": os.Getenv("APP_VERSION"),
			"env":     os.Getenv("APP_ENV"),
		})
	})

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "/api"
	}

	api := app.Group(baseURL)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
			"database":  "connected",
			"redis":     "connected",
		})

	})

	return app
}
