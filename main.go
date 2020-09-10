package main

import (
	"context"
	"net/http"

	"github.com/Mockturnal/voting-app-backend/api/auth"
	"github.com/Mockturnal/voting-app-backend/api/poll"
	"github.com/Mockturnal/voting-app-backend/api/user"
	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	database.Init()

	if err := database.Ping(context.TODO()); err != nil {
		panic(err)
	}

	if err := database.Migrate(&user.User{}, &poll.Poll{}); err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	usersGroup := e.Group("/users")
	pollsGroup := e.Group("/polls")
	authGroup := e.Group("/auth")

	user.NewUserRoutes(usersGroup)
	poll.NewPollRoutes(pollsGroup)
	auth.NewAuthRoutes(authGroup)

	e.Logger.Fatal(e.Start(":5000"))
}
