package database

import (
	"api/config"
	"api/middlewares"
	"api/models"
)

func LoginRepository(data *models.User) (*models.User, error) {
	var err error
	DB := config.DbManager()

	if err = DB.Where("email = ? AND password = ?", data.Email, data.Password).First(&data).Error; err != nil {
		return nil, err
	}

	if data.Token, err = middlewares.CreateToken(data.Role); err != nil {
		return nil, err
	}

	return data, nil
}

func RegisterRepository(data *models.User) (*models.User, error) {
	var err error
	DB := config.DbManager()

	if err := DB.Create(&data).Error; err != nil {
		return nil, err
	}

	if err = DB.Save(data).Error; err != nil {
		return nil, err
	}

	if data.Token, err = middlewares.CreateToken(data.Role); err != nil {
		return nil, err
	}

	return data, nil
}
