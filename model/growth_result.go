package model

import "time"
type GrowthResult struct {
	ID         int       `json:"id"`
	Result  string `json:"result"`
	CreatedAt  time.Time `json:"created_at"`
}