package service

import (

	//"net/http"
	"github.com/kirankkirankumar/gqlgen-ddk/repository"
)

type Service interface {
	//ParseFile(r *http.Request) (*model.Schema, error)
	//FileSave(r *http.Request) (string, string, error)
	GenerateModel() error
	//Middleware() func(http.Handler) http.Handler
	MigrateModels()
	//CtxValue(ctx context.Context) *model.Claims
}

type service struct {
	Repo    repository.Repository
	
}

func NewService(repo repository.Repository) (Service, error) {

	return &service{Repo: repo}, nil
}