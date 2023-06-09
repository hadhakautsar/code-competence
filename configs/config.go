package config

import (
	"code-competence-remidi/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/db_barang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&models.Barang{})
	if err != nil {
		return nil, fmt.Errorf("failed to automigrate models: %w", err)
	}

	return db, nil
}
