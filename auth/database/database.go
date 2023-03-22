package database

import (
	"log"
	"os"

	"github.com/dodysat/gue-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	db, err := gorm.Open(mysql.Open(dbUser+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database\n", err.Error())
		os.Exit(2)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.User{}, &models.Session{}, &models.Activity{})
	Database = DbInstance{Db: db}
}
