package auth

import "github.com/labstack/echo"

func NewAuthRoutes(g *echo.Group) {
	g.POST("/login", Login)
	g.POST("/register", Register)
}
