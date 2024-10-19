package main

import (
	"fmt"
	"go-practice/model"
	"go-practice/router"
	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get *sql.DB from GORM:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Connected to MySQL using GORM!")

	r := router.SetupRouter(db)

	err = db.AutoMigrate(&model.Board{})
	if err != nil {
		log.Fatal("Failed to migrate table:", err)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
