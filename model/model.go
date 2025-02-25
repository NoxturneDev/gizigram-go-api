package model

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	*gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	Parent    *Parent `json:"parent" gorm:"foreignKey:UserID;references:ID"`
}

type Admin struct {
	*gorm.Model
	Name      string    `json:"name"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Children struct {
	*gorm.Model
	Name         string         `json:"name"`
	Age          int            `json:"age"`
	Weight       int            `json:"weight"`
	Height       int            `json:"height"`
	Gender       int            `json:"gender"`
	ParentID     int            `json:"parent_id"`
	Parent       Parent         `json:"parent" gorm:"foreignKey:ParentID;references:ID"`
	BirthDate    time.Time      `json:"birth_date"`
	CreatedAt    time.Time      `json:"created_at"`
	GrowthRecord []GrowthRecord `json:"growth_record" gorm:"foreignKey:ChildrenID;references:ID"`
}

type Parent struct {
	*gorm.Model
	Name      string     `json:"name"`
	UserID    int        `json:"user_id"`
	Height    int        `json:"height"`
	Address   string     `json:"address"`
	Gender    int        `json:"gender"`
	NIK       string     `json:"nik"`
	CreatedAt time.Time  `json:"created_at"`
	Children  []Children `json:"children" gorm:"foreignKey:ParentID;references:ID"`
}

type GrowthRecord struct {
	*gorm.Model
	RecordCount   int           `json:"record_count"`
	ChildrenID    int           `json:"children_id"`
	Children      Children      `json:"children" gorm:"foreignKey:ChildrenID;references:ID"`
	WeightAfter   int           `json:"weight_after"`
	WeightBefore  int           `json:"weight_before"`
	HeightAfter   int           `json:"height_after"`
	HeightBefore  int           `json:"height_before"`
	AddedHeight   int           `json:"added_height"`
	AddedWeight   int           `json:"added_weight"`
	LastCheckDate time.Time     `json:"last_check_date"`
	RecordDate    time.Time     `json:"record_date"`
	CreatedAt     time.Time     `json:"created_at"`
	GrowthResult  *GrowthResult `json:"growth_result" gorm:"foreignKey:GrowthRecordID;references:ID"`
}

type GrowthResult struct {
	*gorm.Model
	Result         string       `json:"result"`
	GrowthRecord   GrowthRecord `json:"growth_record" gorm:"foreignKey:GrowthRecordID;references:ID"`
	GrowthRecordID int          `json:"growth_record_id"`
}
