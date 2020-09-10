package routes

import (
	"net/http"

	"github.com/Mockturnal/voting-app-backend/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To Mockturnal API")
	})

	// Endpoint User
	e.GET("/users", handlers.GetUsers)
	e.POST("/users/register", handlers.Register)

	e.GET("/polls", handlers.GetPolls)
	e.POST("/polls", handlers.CreatePolls)

	return e
}
