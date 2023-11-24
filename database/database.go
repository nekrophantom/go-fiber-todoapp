package database

import (
	"log"

	"github.com/nekrophantom/go-fiber-todoapp/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConnection()  {

	dialect := mysql.Open("root:@tcp(localhost:3306)/go_fiber_todoapp?charset=utf8mb4&parseTime=True&loc=Local")
	
	var err error
	DB, err = gorm.Open(dialect, &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed connect to the database!", err)
	}

	log.Println("Connected to the database successfully")
	log.Println("Running Migrations")
	
	// Todo : Migration 
	AutoMigrateModels()

}

func AutoMigrateModels() {
	DB.AutoMigrate(&models.Todo{})
}