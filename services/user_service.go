package services

import (
	"github.com/berkatps/database"
	"github.com/berkatps/model"
)

func CreateUser(user *model.Users) error {
	return database.DB.Create(&user).Error
}

func GetUser() ([]model.Users, error) {
	var users []model.Users
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteUser(id int) error {
	return database.DB.Delete(&model.Users{}, id).Error
}
