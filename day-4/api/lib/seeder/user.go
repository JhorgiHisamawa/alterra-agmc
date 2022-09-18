package user

import (
	"api/config"
	"api/lib/database"
	"api/models"
)

func NewUserSeeder() error {

	seeds := []models.User{
		{
			Id:       1,
			Name:     "testing1",
			Email:    "testing@gmail.com",
			Password: "passwordtest",
			Role:     "test",
			Token:    "",
		},
		{
			Id:       2,
			Name:     "testing2",
			Email:    "testing@gmail.com",
			Password: "passwordtest",
			Role:     "test",
			Token:    "",
		},
		{
			Id:       3,
			Name:     "testing3",
			Email:    "testing@gmail.com",
			Password: "passwordtest",
			Role:     "test",
			Token:    "",
		},
	}
	for _, seed := range seeds {
		_, err := database.CreateUser(seed)
		return err
	}
	return nil
}

func DeleteSeeder() {
	DB := config.DbManager()
	DB.Exec("DELETE FROM users")
}
