package auth

import (
	"net/http"
	"time"

	"github.com/Mockturnal/voting-app-backend/api/user"
	"github.com/Mockturnal/voting-app-backend/auth/jwt"
	"github.com/Mockturnal/voting-app-backend/config"
	"github.com/Mockturnal/voting-app-backend/database"
	"github.com/Mockturnal/voting-app-backend/helpers"
	gojwt "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	var (
		v        = validator.New()
		email    = c.FormValue("email")
		password = c.FormValue("password")
		cfg      = config.GetConfig()
	)

	db := database.GetConnection()

	user := new(user.User)

	if err := v.Struct(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err := db.Model(user).Where("email = ?", email).First()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.Response{
			Status:  http.StatusUnauthorized,
			Message: "Invalid email / password",
		})
	}

	matched := helpers.HashCompare(user.Password, password)
	if !matched {
		return c.JSON(http.StatusUnauthorized, helpers.Response{
			Status:  http.StatusUnauthorized,
			Message: "Invalid email / password",
		})
	}

	at := &jwt.Token{
		JwtAlgo:   gojwt.SigningMethodHS256,
		JwtClaim:  jwt.NewClaim(user.ID, time.Now().Add(time.Minute*10).Unix()),
		JwtSecret: cfg.Server.AccessTokenSecret,
	}

	rt := &jwt.Token{
		JwtAlgo:   gojwt.SigningMethodHS256,
		JwtClaim:  jwt.NewClaim(user.ID, time.Now().Add(time.Hour*24*7).Unix()),
		JwtSecret: cfg.Server.RefreshTokenSecret,
	}

	accessToken, err := at.GetToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	refreshToken, err := rt.GetToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func Register(c echo.Context) error {
	var (
		v        = validator.New()
		username = c.FormValue("username")
		email    = c.FormValue("email")
		password = c.FormValue("password")
	)
	db := database.GetConnection()
	hash, _ := helpers.GenerateHash(password)

	data := new(user.User)
	data.Username = username
	data.Email = email
	data.Password = hash
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	if err := v.Struct(data); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	_, err := db.Model(data).Insert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Successfully registered",
	})
}
