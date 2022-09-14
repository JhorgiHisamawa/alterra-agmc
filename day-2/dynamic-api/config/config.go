package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := Config("DB_USER") + ":" + Config("DB_PASSWORD") + "@tcp(" + Config("DB_HOST") + ":" + Config("DB_PORT") + ")/" + Config("DB_NAME") + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}
	DB.AutoMigrate()
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
