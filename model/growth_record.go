package model

import "time"

type GrowthRecord struct {
	ID             int       `json:"id"`
	RecordCount    int       `json:"record_count"`
	ChildrenID     int       `json:"children_id"`
	GrowthResultID int       `json:"growth_result_id"`
	WeightAfter    int       `json:"weight_after"`
	WeightBefore   int       `json:"weight_before"`
	HeightAfter    int       `json:"height_after"`
	HeightBefore   int       `json:"height_before"`
	AddedHeight    int       `json:"added_height"`
	AddedWeight    int       `json:"added_weight"`
	CreatedAt      time.Time `json:"created_at"`
}
