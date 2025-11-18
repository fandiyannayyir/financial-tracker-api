package main

import (
	"financial-tracker-api/config"
	"financial-tracker-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Create a new Echo instance
	e := echo.New()
	routes.RegisterRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
