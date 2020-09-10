package poll

import (
	"time"

	"github.com/Mockturnal/voting-app-backend/api/user"
)

type PollOption struct {
	Option string      `json:"option"`
	Users  []user.User `json:"users"`
}

type Poll struct {
	tableName struct{} `pg:"polls"`

	ID      uint64       `pg:",pk" json:"id"`
	Title   string       `json:"title"`
	Options []PollOption `json:"options"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
