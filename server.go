package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	//"time"
	"context"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kirankkirankumar/gqlgen-ddk/graph"
	"github.com/kirankkirankumar/gqlgen-ddk/graph/generated"
	"github.com/kirankkirankumar/gqlgen-ddk/repository"
	"github.com/kirankkirankumar/gqlgen-ddk/service"
	"github.com/go-chi/chi"
	"github.com/kirankkirankumar/gqlgen-ddk/handler"
	"github.com/99designs/gqlgen/graphql"
)

const defaultPort = "8089"



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
    
	//started
	fmt.Printf("Welcome")
	s, err := service.NewService(repo)

	//s.GenerateModel()
	s.MigrateModels()

	h := &handler.Handler{
		Repo:    repo,
		Service: s}

	c := generated.Config{Resolvers: &graph.Resolver{Repo: repo}}

    c.Directives.Mapping = func(ctx context.Context, obj interface{}, next graphql.Resolver, typeArg *string) (interface{}, error) {

      return next(ctx)

   }



    srv := gqlHandler.NewDefaultServer(

        generated.NewExecutableSchema(c))
	
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		
		r.Handle("/", playground.Handler("GraphQL playground", "/query"))
		r.Handle("/query", srv)
		r.HandleFunc("/upload-schema", h.WriteFile)
	})
	
	
	
	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
