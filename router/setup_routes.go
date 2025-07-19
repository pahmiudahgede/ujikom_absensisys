package router

import (
	"absensibe/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group(os.Getenv("BASE_URL"))
	api.Use(middleware.APIKeyValidator())
	AuthRouter(api)
	StudentRoutes(api)
}
