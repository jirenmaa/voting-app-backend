package poll

import (
	"fmt"
	"time"
	"net/http"
	"reflect"

	"github.com/Mockturnal/voting-app-backend/api/user"
	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/labstack/echo"
)

// TempData godoc
// @description keeping temporary data
type TempData struct {
	option string
	user   []interface{}
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

	db := database.GetConnection()
	data := new(Poll)
	dataPoll := []PollOption{}

	for i := 1; i < 3; i++ {
		optName := fmt.Sprintf("Test %d", i)
		temp := PollOption{
			Option: optName,
			Users:  []user.User{},
		}

		dataPoll = append(dataPoll, temp)
	}

	data.Title = res["Title"].(string)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.Options = dataPoll

	result, err := db.Model(data).Insert()
	if err != nil && result != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(reflect.TypeOf(dataPoll))
	return c.JSON(http.StatusOK, data)
	// return c.JSON(http.StatusOK, res)
}
