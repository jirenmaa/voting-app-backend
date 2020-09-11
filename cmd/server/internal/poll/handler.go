package poll

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type PollController interface {
	GetPolls(ctx *gin.Context)
	CreatePoll(ctx *gin.Context)
	DeletePoll(ctx *gin.Context)
}

// PollServices ...
type PollServices struct {
	Db *pg.DB
}

// NewPollService ...
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
	if err := u.Db.Model(data).Column("id", "title", "options", "created_at", "updated_at").Select(); err != nil {
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
func (u *PollServices) CreatePoll(ctx *gin.Context) {
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
		"message": "create poll success",
		"data":    poll,
	})
}

// @Description Delete poll
// @Accept json
// @Produce json
// @Success 200 {object} []Poll
// @Failure 500 {object} gin.H
// @Router /polls [delete]
func (u *PollServices) DeletePoll(ctx *gin.Context) {
	poll := Poll{}
	id, _ := strconv.Atoi(ctx.Query("id"))

	_, err := u.Db.Model(&poll).Where("id = ?", id).Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete poll success",
		"data":    []interface{}{},
	})
}
