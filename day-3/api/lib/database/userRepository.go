package database

import (
	"api/config"
	"api/models"
)

func CreateUser(data models.User) (*models.User, error) {
	DB := config.DbManager()
	if err := DB.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func GetUser() (data *[]models.User, err error) {
	DB := config.DbManager()
	if err := DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func GetUserByID(id int) (data *models.User, err error) {
	DB := config.DbManager()
	if err := DB.Where(`id=?`, id).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func UpdateUser(id int, data models.User) (*models.User, error) {
	DB := config.DbManager()
	if err := DB.Where(`id=?`, id).Updates(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func DeleteUser(id int) (data *models.User, err error) {
	DB := config.DbManager()
	if err := DB.Where(`id=?`, id).Delete(&data, id).Error; err != nil {
		return nil, err
	}

	return data, nil
}
