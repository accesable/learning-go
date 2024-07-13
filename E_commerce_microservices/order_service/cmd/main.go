package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"trann/ecom/order_service/cmd/api"
)

var db *gorm.DB

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize GORM
	db, err = initGorm()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func main() {
	srv := api.NewAPIServer(":8087", nil)
	if err := srv.Run(); err != nil {
		log.Fatal("Unable to Start up Server")
	}
}

// initGorm initializes the GORM database connection
func initGorm() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	// DSN (Data Source Name) format for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
