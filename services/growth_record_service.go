package services

import (
	"fmt"
	"gizigram-go-api/model"
	"gorm.io/gorm"
)

func CreateGrowthRecord(tx *gorm.DB, children *model.Children) error {
	var existingChild model.Children
	err := tx.Where("id = ?", children.ID).First(&existingChild).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("child with ID %d not found", children.ID)
		}
		return err
	}

	var lastGrowthRecord model.GrowthRecord
	err = tx.Where("children_id = ?", children.ID).Order("created_at desc").First(&lastGrowthRecord).Error
	isNew := err == gorm.ErrRecordNotFound

	growthRecord := model.GrowthRecord{
		ChildrenID:  children.ID,
		WeightAfter: children.Weight,
		HeightAfter: children.Height,
		RecordCount: lastGrowthRecord.RecordCount + 1,
	}

	if isNew {
		growthRecord.RecordCount = 0
		growthRecord.WeightBefore = 0
		growthRecord.HeightBefore = 0
		growthRecord.AddedWeight = children.Weight
		growthRecord.AddedHeight = children.Height

		if err := tx.Create(&growthRecord).Error; err != nil {
			return err
		}

		if err := tx.Create(&children).Error; err != nil {
			return err
		}
	} else {
		growthRecord.WeightBefore = lastGrowthRecord.WeightAfter
		growthRecord.HeightBefore = lastGrowthRecord.HeightAfter
		growthRecord.AddedWeight = children.Weight - lastGrowthRecord.WeightAfter
		growthRecord.AddedHeight = children.Height - lastGrowthRecord.HeightAfter

		if err := tx.Create(&growthRecord).Error; err != nil {
			return err
		}

		if err := tx.Model(&existingChild).Updates(&children).Error; err != nil {
			return err
		}
	}

	return nil
}
