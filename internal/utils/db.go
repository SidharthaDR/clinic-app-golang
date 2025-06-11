package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/SidharthaDR/golang-clinic-app/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDatabase() {
	LoadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db

	// Migrate User model
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Failed to migrate User model: ", err)
	}

	// Migrt Patient model
	err = db.AutoMigrate(&models.Patient{})
	if err != nil {
		log.Fatal("❌ Failed to migrate Patient model: ", err)
	}

	fmt.Println("✅ Database connected")

}
