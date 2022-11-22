package service

import (

	//"net/http"
	"github.com/kirankkirankumar/gqlgen-ddk/repository"
)

type Service interface {
	
	GenerateModel() error
	MigrateModels()
	writeData() error
}

type service struct {
	Repo    repository.Repository
	
}

func NewService(repo repository.Repository) (Service, error) {

	return &service{Repo: repo}, nil
}

