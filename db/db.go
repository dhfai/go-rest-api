package db

import (
	"fmt"
	"log"

	"github.com/dhfai/go-rest-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBStorage struct {
	DB *gorm.DB
}

func NewDBStorage() (*DBStorage, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Envs.Host, config.Envs.User, config.Envs.Password,
		config.Envs.DBName, config.Envs.Port, config.Envs.SSLMode, config.Envs.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL with GORM!")
	return &DBStorage{DB: db}, nil
}
