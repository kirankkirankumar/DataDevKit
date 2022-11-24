package repository

import (
	"fmt"

	"github.com/kirankkirankumar/gqlgen-ddk/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	GetData(model interface{}) error

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
func migrateSchemaOnLoad(db *gorm.DB) {
	db.AutoMigrate(&model.Schema_Entry{})

	
}

func (r *repository) GetData(model interface{}) error {

	// for _, tables := range preLoad {
	// 	db.Preload(tables)
	// }
	result := r.db.Preload(clause.Associations).Find(model)

	return result.Error
}

func (r *repository) createData(model interface{}) error {

	
	result := r.db.Create(model)

	return result.Error
}