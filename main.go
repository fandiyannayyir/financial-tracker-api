package main

import (
	"financial-tracker-api/config"
	"financial-tracker-api/routes"
	"financial-tracker-api/utils"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Connect to the database
	config.ConnectDB()

	// Create a new Echo instance
	e := echo.New()
	routes.RegisterRoutes(e)

	// Get server port from environment variables
	port := utils.GetEnv("SERVER_PORT", "8080")

	// Start the server
	e.Logger.Fatal(e.Start(":" + port))
}
