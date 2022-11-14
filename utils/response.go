package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, status int, data interface{}, key string) error {
	wrapper := make(map[string]interface{})

	wrapper[key] = data

	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "Application/JSON")
	w.WriteHeader(status)
	n, err := w.Write(json)
	// n, err := fmt.Fprintf(w, string(json))
	if err != nil && n == 0 {
		return err
	}

	return nil
}

func WriteErrorResponse(w http.ResponseWriter, err error) {

	type errjson struct {
		Message string `json:"message"`
	}
	theError := errjson{
		Message: err.Error(),
	}
	WriteJsonResponse(w, http.StatusBadRequest, theError, "error")
}
