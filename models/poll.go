package models

import "time"

type PollOption struct {
	Option string `json:"option"`
	Users  []User `json:"users"`
}

type Poll struct {
	tableName struct{} `pg:"polls"`

	ID      uint16       `pg:",pk" json:"id"`
	Title   string       `json:"title"`
	Options []PollOption `json:"options"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
