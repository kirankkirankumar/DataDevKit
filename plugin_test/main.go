package main

import (
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/kirankkirankumar/gqlgen-ddk/plugins"
)

func main() {
	log.Println("Running")
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	err = api.Generate(cfg, api.AddPlugin(plugins.New()))
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
}