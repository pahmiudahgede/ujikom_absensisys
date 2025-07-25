package main

import (
	"absensibe/config"
	"absensibe/router"
	"log"
	"os"
)

//	@title			UJIKOM API
//	@version		1.0
//	@description	Ini adalah API untuk ujikom mif 2025
//	@termsOfService	http://swagger.io/terms/

func main() {
	config.InitializeAll()

	app := config.SetupServer()
	router.SetupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "7080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📍 API Base URL: %s", os.Getenv("BASE_URL"))
	log.Printf("🌍 Environment: %s", os.Getenv("APP_ENV"))

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
