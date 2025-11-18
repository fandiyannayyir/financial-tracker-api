package routes

import (
	"financial-tracker-api/modules/users"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	users.UserRoutes(api.Group("/users"))
}