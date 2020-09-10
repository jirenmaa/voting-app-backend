package user

import (
	"net/http"

	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := database.GetConnection()
	data := new([]User)
	if err := db.Model(data).Column("username", "email", "created_at", "updated_at").Select(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
