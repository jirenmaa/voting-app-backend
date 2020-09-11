package poll

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type PollController interface {
	GetPolls(ctx *gin.Context)
	CreatePolls(ctx *gin.Context)
}

type PollServices struct {
	Db *pg.DB
}

func NewPollService(db *pg.DB) PollController {
	return &PollServices{
		Db: db,
	}
}

// @Description Poll list
// @Accept json
// @Produce json
// @Success 200 {object} []Poll
// @Failure 500 {object} gin.H
// @Router /polls [get]
func (u *PollServices) GetPolls(ctx *gin.Context) {
	data := new([]Poll)
	if err := u.Db.Model(data).Column("title", "options", "created_at", "updated_at").Select(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get polls success",
		"data":    data,
	})
}

// @Description Create poll
// @Accept json
// @Produce json
// @Success 200 {object} []Poll
// @Failure 500 {object} gin.H
// @Router /polls [post]
func (u *PollServices) CreatePolls(ctx *gin.Context) {
	var request Poll
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	poll := Poll{
		Title:     request.Title,
		Options:   request.Options,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := u.Db.Model(&poll).Insert()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "post polls success",
		"data":    poll,
	})
}
