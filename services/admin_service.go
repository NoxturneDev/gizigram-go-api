package services

import (
	"gizigram-go-api/database"
	"gizigram-go-api/model"
	"gorm.io/gorm"
	"time"
)

func CreateParent(parent *model.Parent, phoneNumber string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		user := model.Users{
			Username: phoneNumber,
			Password: phoneNumber,
		}

		err := tx.Create(&user).Error
		if err != nil {
			return err
		}

		parent.UserID = int(user.ID)
		err = tx.Create(&parent).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ShowAllParrent() ([]model.Parent, error) {
	var parents []model.Parent
	if err := database.DB.Preload("Children").Find(&parents).Error; err != nil {
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

func ShowParentOptions() (interface{}, error) {
	type Options struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	}

	var parents []model.Parent
	var options []Options

	if err := database.DB.Find(&parents).Error; err != nil {
		return nil, err
	}

	for _, parent := range parents {
		options = append(options, Options{
			Value: int(parent.ID),
			Label: parent.Name,
		})
	}

	return options, nil
}

func CreateChildren(tx *gorm.DB, children *model.Children) error {
	if err := tx.Create(children).Error; err != nil {
		return err
	}

	growth := model.GrowthRecord{
		RecordCount:  1,
		ChildrenID:   int(children.ID),
		WeightAfter:  children.Weight,
		WeightBefore: children.Weight,
		HeightAfter:  children.Height,
		HeightBefore: children.Height,
		AddedHeight:  0,
		AddedWeight:  0,
		CreatedAt:    time.Now(),
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
	if err := database.DB.Joins("Parent").Find(&children).Error; err != nil {
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
