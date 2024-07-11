package model

import "time"

type Parent struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	UserID    int    `json:"user_id"`
	CreatedAt time.Time
}
