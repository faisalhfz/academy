package config

import (
	"os"

	"gorm.io/gorm"

	"user-management/entity"

	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&entity.User{})
}
