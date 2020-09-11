package main

import (
	"log"
	"os"
	"time"

	_ "github.com/Mockturnal/voting-app-backend/cmd/server/docs"
	"github.com/Mockturnal/voting-app-backend/cmd/server/internal/auth"
	"github.com/Mockturnal/voting-app-backend/cmd/server/internal/poll"
	"github.com/Mockturnal/voting-app-backend/cmd/server/internal/user"
	"github.com/Mockturnal/voting-app-backend/pkg/cache"
	"github.com/Mockturnal/voting-app-backend/pkg/database"
	"github.com/Mockturnal/voting-app-backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func LoadEnv() error {
	switch os.Getenv("APP_ENV") {
	case "production":
		err := godotenv.Load("config/prod.env")
		return err
	default:
		err := godotenv.Load("config/dev.env")
		return err
	}
}

// @title Voting App API Documentation
// @description This is the backend service documentation for the voting app
// @BasePath /
func main() {
	if err := LoadEnv(); err != nil {
		panic(err)
	}

	conn, err := database.NewConnection(&pg.Options{
		Network:  os.Getenv("POSTGRESQL_NETWORK"),
		Addr:     os.Getenv("POSTGRESQL_ADDR"),
		User:     os.Getenv("POSTGRESQL_USER"),
		Password: os.Getenv("POSTGRESQL_PASSWORD"),
		Database: os.Getenv("POSTGRESQL_DATABASE"),
	})
	if err != nil {
		panic(err)
	}

	if err := conn.CreateSchema(&user.User{}, &poll.Poll{}); err != nil {
		panic(err)
	}

	jwtAuthService := jwt.NewJWTService()
	redisService := cache.NewRedisCache(os.Getenv("REDIS_ADDR"), 0, 2*time.Hour)
	authController := auth.NewAuthService(conn.DB, redisService, jwtAuthService)
	userController := user.NewUserService(conn.DB)
	pollController := poll.NewPollService(conn.DB)

	r := gin.Default()

	users := r.Group("/users")
	{
		users.GET("/", userController.GetUsers)
		users.DELETE("/:id", userController.DelUsers)
	}

	polls := r.Group("/polls")
	{
		polls.GET("/", pollController.GetPolls)
		polls.POST("/", pollController.CreatePoll)
		polls.DELETE("/", pollController.DeletePoll)
		polls.PUT("/", pollController.UpdatePoll)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

	swaggerURL := ginSwagger.URL("http://localhost:5000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	log.Fatal(r.Run(":5000"))
}
