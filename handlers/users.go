package handlers

import (
	"net/http"

	helper "github.com/Mockturnal/voting-app-backend/helpers"

	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/Mockturnal/voting-app-backend/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := database.GetConnection()
	data := new(models.User)
	err := db.Model(data).Select()

	return c.JSON(http.StatusOK, err)
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
	err := v.Struct(data)
	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}
	db.Model(data).Insert()
	return c.JSON(http.StatusOK, data)
}
