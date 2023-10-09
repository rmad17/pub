package models

import (
    "fmt"
    "log"
    "os"

	"gorm.io/gorm"
    "gorm.io/driver/postgres"
    "github.com/joho/godotenv"
)
var db *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}	
	
	// Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	// DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	// db, err = gorm.Open(Dbdriver, DBURL)
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", DbHost, DbUser, DbPassword, DbName, DbPort)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})	

    
    if err != nil {
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", err)
	} else {
        fmt.Println("We are connected to the database: ", db)
	}
    
    db.AutoMigrate(&User{})
}
