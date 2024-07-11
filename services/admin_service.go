package services

import (
	"github.com/berkatps/database"
	"github.com/berkatps/model"
	"gorm.io/gorm"
	"time"
)

func CreateParent(parent *model.Parent) error {
	return database.DB.Create(&parent).Error
}

func ShowAllParrent() ([]model.Parent, error) {
	var parents []model.Parent
	if err := database.DB.Find(&parents).Error; err != nil {
		return nil, err
	}
	return parents, nil
}

func ShowParrentByID(id int) (*model.Parent, error) {
	var parent model.Parent
	if err := database.DB.First(&parent, id).Error; err != nil {
		return nil, err
	}
	return &parent, nil
}

func CreateChildren(tx *gorm.DB, children *model.Children) error {
	if err := tx.Create(children).Error; err != nil {
		return err
	}

	growth := model.GrowthRecord{
		RecordCount:    1,
		ChildrenID:     children.ID,
		GrowthResultID: 0,
		WeightAfter:    children.Weight,
		WeightBefore:   children.Weight,
		HeightAfter:    children.Height,
		HeightBefore:   children.Height,
		AddedHeight:    0,
		AddedWeight:    0,
		CreatedAt:      time.Now(),
	}

	if err := tx.Create(&growth).Error; err != nil {
		return err
	}

	return nil
}

func GetChildrenMatchByParentID(id int) ([]model.Children, error) {
	var children []model.Children
	if err := database.DB.Where("parent_id = ?", id).Find(&children).Error; err != nil {
		return nil, err
	}
	return children, nil
}

func ShowAllChildren() ([]model.Children, error) {
	var children []model.Children
	if err := database.DB.Find(&children).Error; err != nil {
		return nil, err
	}
	return children, nil
}

func ShowChildrenByID(id int) (*model.Children, error) {
	var children model.Children
	if err := database.DB.First(&children, id).Error; err != nil {
		return nil, err
	}
	return &children, nil
}

func DeleteChildren(id int) error {
	return database.DB.Delete(&model.Children{}, id).Error
}

func DeleteParent(id int) error {
	return database.DB.Delete(&model.Parent{}, id).Error
}
