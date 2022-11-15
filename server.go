package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kirankkirankumar/gqlgen-ddk/graph"
	"github.com/kirankkirankumar/gqlgen-ddk/graph/generated"
	"github.com/kirankkirankumar/gqlgen-ddk/repository"
	"github.com/kirankkirankumar/gqlgen-ddk/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := &repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}

	repo, err := repository.NewRepository(config)
	if err != nil {
		log.Fatal("could not connect to db")
	}
    
	fmt.Printf("Welcome")
	s, err := service.NewService(repo)

	s.MigrateModels()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
