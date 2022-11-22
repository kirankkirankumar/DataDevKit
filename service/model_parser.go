package service

import (
	"fmt"
	"log"
	"github.com/kirankkirankumar/gqlgen-ddk/graph/model"
)

func (s *service) MigrateModels() {

	

	structs := model.GetStructs()

	for names, structss := range structs {

		
		 err := s.Repo.MigrateSchema(names, structss)
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Println("Migrated Successfully..")

	return

}


