package middlewares

import (
	"api/config"
	"api/models"

	"github.com/labstack/echo"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	DB := config.DbManager()
	var data models.User
	var err error

	if err = DB.Where("email = ? AND password = ?", email, password).First(&data).Error; err != nil {
		return false, err
	}

	if data.Role != ADMIN {
		return false, err
	}
	return true, nil
}
