package services

import (
	"github.com/berkatps/database"
	"github.com/berkatps/model"
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

func CreateChildren(children *model.Children) error {
	return database.DB.Create(&children).Error
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
