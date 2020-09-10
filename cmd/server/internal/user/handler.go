package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type UserController interface {
	GetUsers(ctx *gin.Context)
}

type UserServices struct {
	Db *pg.DB
}

func NewUserService(db *pg.DB) UserController {
	return &UserServices{
		Db: db,
	}
}

// @Description Users list
// @Accept json
// @Produce json
// @Success 200 {object} []User
// @Failure 500 {object} gin.H
// @Router /users [get]
func (u *UserServices) GetUsers(ctx *gin.Context) {
	data := new([]User)
	if err := u.Db.Model(data).Column("username", "email", "created_at", "updated_at").Select(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get users success",
		"data":    data,
	})
}
