package database

import (
	"log"
	"os"

	"github.com/dodysat/gue-order/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(mysql.Open("root:@tcp(order-database:3306)/order?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	// db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3436)/order?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database\n", err.Error())
		os.Exit(2)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.Cart{}, &models.Checkout{})
	Database = DbInstance{Db: db}
}
