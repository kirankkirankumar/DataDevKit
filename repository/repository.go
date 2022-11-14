package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Repository interface {
	//VideoRepository
	SchemaRepository
}

type repository struct {
	db *gorm.DB
}

func NewRepository(config *Config) (Repository, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	r := &repository{
		db: db,
	}

	return r, nil

}
