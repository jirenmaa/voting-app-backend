package handlers

import (
	"fmt"
	"net/http"
	"time"

	helper "github.com/Mockturnal/voting-app-backend/helpers"

	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/Mockturnal/voting-app-backend/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := database.GetConnection()
	data := new([]models.User)
	if err := db.Model(data).Column("username", "email", "created_at", "updated_at").Select(); err != nil {
		fmt.Print(data)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func Register(c echo.Context) error {
	var (
		v        = validator.New()
		username = c.FormValue("username")
		email    = c.FormValue("email")
		password = c.FormValue("password")
	)
	db := database.GetConnection()
	hash, _ := helper.HashPassword(password)
	data := new(models.User)
	data.Username = username
	data.Email = email
	data.Password = hash
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	if err := v.Struct(data); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := db.Model(data).Insert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
