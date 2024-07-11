package model

import "time"

type Users struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
}
