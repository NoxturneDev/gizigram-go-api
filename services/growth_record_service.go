package services

import (
	"fmt"
	"gizigram-go-api/database"
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
		ChildrenID:  int(children.ID),
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

func CreateGrowthRecordWithoutChildren(growthRecord *model.GrowthRecord) error {
	var children model.Children

	fmt.Println(growthRecord)
	err := database.DB.Where("id = ?", growthRecord.ChildrenID).First(&children).Error
	if err != nil {
		return err
	}

	var lastGrowthRecord model.GrowthRecord
	err = database.DB.Where("children_id = ?", children.ID).Order("created_at desc").First(&lastGrowthRecord).Error
	if err != nil {
		return err
	}

	growthRecord.WeightBefore = children.Weight
	growthRecord.HeightBefore = children.Height
	growthRecord.LastCheckDate = lastGrowthRecord.CreatedAt
	growthRecord.AddedWeight = growthRecord.WeightAfter - growthRecord.WeightBefore
	growthRecord.AddedHeight = growthRecord.HeightAfter - growthRecord.HeightBefore
	growthRecord.RecordCount = lastGrowthRecord.RecordCount + 1

	children.Weight = growthRecord.WeightAfter
	children.Height = growthRecord.HeightAfter

	if err := database.DB.Create(&growthRecord).Error; err != nil {
		return err
	}

	// update children with new weight and height
	if err := database.DB.Model(&children).Updates(&children).Error; err != nil {
		return err
	}

	return nil
}

func ShowGrowthRecordByChildrenID(id int) ([]model.GrowthRecord, error) {
	var growthRecords []model.GrowthRecord

	if err := database.DB.Where("children_id = ?", id).Find(&growthRecords).Error; err != nil {
		return nil, err
	}

	return growthRecords, nil
}
