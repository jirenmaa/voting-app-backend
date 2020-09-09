package user

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewUserRoutes(g *echo.Group) {
	g.GET("", GetUsers, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(os.Getenv("ACCESS_TOKEN_SECRET")),
		SigningMethod: jwt.SigningMethodHS256.Name,
	}))
}
