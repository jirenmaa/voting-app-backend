package handlers

import (
	"net/http"

	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/Mockturnal/voting-app-backend/models"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := database.GetConnection()
	data := new(models.User)
	err := db.Model(data).Select()

	return c.JSON(http.StatusOK, err)
}
