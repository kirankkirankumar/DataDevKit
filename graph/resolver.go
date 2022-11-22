package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go get -d github.com/99designs/gqlgen
//go:generate go run github.com/kirankkirankumar/gqlgen-ddk/plugin_test




import "github.com/kirankkirankumar/gqlgen-ddk/repository"

type Resolver struct {
	Repo repository.Repository
}