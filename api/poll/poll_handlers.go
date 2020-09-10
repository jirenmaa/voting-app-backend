package poll

import (
	"fmt"
	"time"
	"net/http"
	"reflect"

	// "time"

	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/Mockturnal/voting-app-backend/helpers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

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
	db := database.GetConnection()
	request := helpers.GetJSON(c)

	data := new(Poll)
	dataPoll := []PollOption{}

	// for i := 1; i < 3; i++ {
	// 	optName := fmt.Sprintf("Test %d", i)
	// 	temp := PollOption{
	// 		Option: optName,
	// 		Users:  []user.User{},
	// 	}

	// 	dataPoll = append(dataPoll, temp)
	// }

	var (
		_optName = ""
	)

	for _, value := range request["Options"].([]interface{}) {
		for _, n := range value.(map[string]interface{}) {
			if reflect.TypeOf(n).Name() == "string" {
				_optName = n.(string)
			}
		}

		temp := PollOption{
			Option: _optName,
		}

		dataPoll = append(dataPoll, temp)
	}

	var v = validator.New()

	data.Title = request["Title"].(string)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.Options = dataPoll

	if err := v.Struct(data); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	_, err := db.Model(data).Insert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
	// return c.JSON(http.StatusOK, dataPoll)
}
