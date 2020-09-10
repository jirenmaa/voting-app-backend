package user

import (
	"time"
)

type User struct {
	tableName struct{} `pg:"users"`

	ID       uint16 `pg:",pk" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
