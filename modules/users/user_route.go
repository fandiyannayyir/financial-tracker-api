package users

import "github.com/labstack/echo/v4"

func UserRoutes(g *echo.Group) {
	repo := NewUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)
}