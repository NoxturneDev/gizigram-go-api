package model

import "time"

type Admin struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
