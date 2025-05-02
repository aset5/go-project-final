package config

import (
	"awesomeProject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=andaset2005 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("AutoMigrate error:", err)
	}
}
