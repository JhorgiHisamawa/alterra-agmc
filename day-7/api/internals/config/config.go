package config

import (
	"api/internals/models"
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type (
	CustomValidator struct {
		Validator *validator.Validate
	}
)

func ConnectDB() {
	dsn := Config("DB_USER") + ":" + Config("DB_PASSWORD") + "@tcp(" + Config("DB_HOST") + ":" + Config("DB_PORT") + ")/" + Config("DB_NAME") + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect database")
	}
	DB.AutoMigrate(&models.User{})
}

func DbManager() *gorm.DB {
	return DB
}

// Config func to get env value from key ---
func Config(key string) string {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
