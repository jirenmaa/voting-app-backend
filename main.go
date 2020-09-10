package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Mockturnal/voting-app-backend/api/auth"
	"github.com/Mockturnal/voting-app-backend/api/poll"
	"github.com/Mockturnal/voting-app-backend/api/user"
	"github.com/Mockturnal/voting-app-backend/config"
	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.LoadConfig()
	cfg := config.GetConfig()
	database.CreateConnection(&pg.Options{
		Addr:     cfg.Database.Addr,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Database: cfg.Database.Db,
	})

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

	addr := fmt.Sprintf("%s:%s",
		cfg.Server.Host,
		cfg.Server.Port)

	e.Logger.Fatal(e.Start(addr))
}
