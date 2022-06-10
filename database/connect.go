package database

import (
	"log"
	"os"
	"simple-blog/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal()
	}
	dsn := os.Getenv("DSN")
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database can't connect")
	} else {
		log.Println("DB connected")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
