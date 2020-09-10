package poll

import (
	"fmt"
	"net/http"

	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/labstack/echo"
)

// TempData godoc
// @description keeping temporary data
type TempData struct {
	option string
	user []interface{}
}

// GetPolls godoc
// @Summary Get all Polls
// @Produce  json
// @Success 200 {object} PollJSONResponse
// @Failure 500 {object} err.Error
// @Router /polls [get]
func GetPolls(c echo.Context) error {
	db := database.GetConnection()
	data := new([]Poll)

	if err := db.Model(data).Column("id", "title", "options", "created_at", "updated_at").Select(); err != nil {
		fmt.Print(data)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

// CreatePolls godoc
// @Summary Create Polls
// @Accept  json
// @Produce  json
// @Param pool body CreatePollRequest true "Creating polls"
// @Success 200 {object} PollJSONResponse
// @Failure 500 {object} err.Error
// @Router /polls [post]
func CreatePolls(c echo.Context) error {
	res := echo.Map{}
	if err := c.Bind(&res); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
