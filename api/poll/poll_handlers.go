package poll

import (
	"fmt"
	"net/http"

	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/labstack/echo"
)

func GetPolls(c echo.Context) error {
	db := database.GetConnection()
	data := new([]Poll)

	if err := db.Model(data).Column("id", "title", "options", "created_at", "updated_at").Select(); err != nil {
		fmt.Print(data)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func CreatePolls(c echo.Context) error {
	return c.String(http.StatusOK, "Yeah")
}
