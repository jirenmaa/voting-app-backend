package helpers

import (
	"github.com/labstack/echo"
)

// GetJSON ...
// @description returning JSON data from echo request body
func GetJSON(e echo.Context) echo.Map {
	result := echo.Map{}
	if err := e.Bind(&result); err != nil {
		panic(err)
	}

	return result
}
