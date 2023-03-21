package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"log"
	"github.com/dodysat/gue-product/models"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb(){
	// db, err := gorm.Open(mysql.Open("root:@tcp(product-database:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3434)/product?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database\n", err.Error())
		os.Exit(2)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.Product{})
	Database = DbInstance{Db: db}
}
