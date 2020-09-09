package routes

import (
	"net/http"

	"github.com/Mockturnal/voting-app-backend/handlers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To Mockturnal API")
	})

	// Endpoint User
	e.GET("/users", handlers.GetUsers)
	e.POST("/users/register", handlers.Register)

	// e.POST("/login", Login)
	// e.POST("/register", SignUp)

	return e
}
