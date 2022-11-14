package repository

import (
	"github.com/kirankkirankumar/gqlgen-ddk/model"
)

type SchemaRepository interface {
	UpdateSchema(entry *model.Schema_Entry) error
	MigrateSchema(name string, table interface{}) error
	ExecQuery(query string) error
}

func (r *repository) UpdateSchema(entry *model.Schema_Entry) error {
	err := r.db.Create(entry).Error
	return err
}

func (r *repository) MigrateSchema(name string, table interface{}) error {

	err := r.db.Table(name).AutoMigrate(table)

	return err
}

func (r *repository) ReloadSchema() error {

	// file,err:=parser.ParseFile("",)

	// file.Package

	// err := r.db.AutoMigrate(&model.Schema_Entry{})

	return nil
}

func (r *repository) ExecQuery(query string) error {

	err := r.db.Exec(query)
	return err.Error
}