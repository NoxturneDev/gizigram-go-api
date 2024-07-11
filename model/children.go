package model

import "time"

type Children struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Height    string `json:"height"`
	ParentID  int    `json:"parent_id"`
	CreatedAt time.Time
}
