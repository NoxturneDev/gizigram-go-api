package model

import "time"

type Children struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Weight    int       `json:"weight"`
	Height    int       `json:"height"`
	Gender    int       `json:"gender"`
	ParentID  int       `json:"parent_id"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
}
