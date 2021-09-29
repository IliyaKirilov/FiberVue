package database

import (
	"fmt"
	//"go-authentication-boilerplate/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file \n", err)
	}

	port := os.Getenv("DB_PORT")
	uname := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
	uname, pass, dbName, port)
	

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.info),
	})

	if err = nil {
		log.Fatal("Failed to connect to PostgresSQL. \n", err)
		os.Exit(2)
	}
	

	fmt.Println("Connected to DB")
}
