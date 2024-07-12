package services

import (
	"fmt"
	"gizigram-go-api/database"
	"gizigram-go-api/model"
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

func GetUserByPhoneNumber(phoneNumber string) *model.Users {
	var user model.Users

	if err := database.DB.Where("username = ?", phoneNumber).First(&user).Error; err != nil {
		return nil
	}

	return &user
}

func DeleteUser(id int) error {
	return database.DB.Delete(&model.Users{}, id).Error
}

func LoginUser(username string, password string) (*model.Users, error) {
	var user *model.Users
	if err := database.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}
	return user, nil
}
