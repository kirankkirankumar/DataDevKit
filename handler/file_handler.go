package handler

import (
	"log"
	"net/http"
	"encoding/json"
	//"io/ioutil"
	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/kirankkirankumar/gqlgen-ddk/model"
	"github.com/kirankkirankumar/gqlgen-ddk/repository"
	"github.com/kirankkirankumar/gqlgen-ddk/service"
	"github.com/kirankkirankumar/gqlgen-ddk/utils"
)

type Handler struct {
	// Schema *graphql.Schema
	Repo    repository.Repository
	Service service.Service
	SRV     *gqlHandler.Server
}

// func (h *Handler) PostFile(w http.ResponseWriter, r *http.Request) {
// 	schema, err := service.ParseFile(r)
// 	if err != nil {
// 		writeErrorResponse(w, errors.New("an error occurred"))
// 		return
// 	}

// 	h.Schema = service.ReloadSchema(schema)
// 	writeJsonResponse(w, http.StatusOK, "file uploaded successfully", "message")
// }



func (h *Handler) WriteFile(w http.ResponseWriter, r *http.Request)  {
	log.Println("hello writing file ")
	log.Println(r.Host)
	log.Println(r)

	var d model.SchemaData

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		utils.WriteJsonResponse(w, http.StatusBadRequest, err.Error(), "message")
		return
	}

	log.Println(d.Data)
	log.Println("hello")
    

	utils.WriteJsonResponse(w, http.StatusOK, "Data written ", "message")
}