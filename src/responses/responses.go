package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON return a json response
func JSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if payload != nil {
		if error := json.NewEncoder(w).Encode(payload); error != nil {
			log.Fatal(error)
		}
	}

}

//Error return a error in json
func Error(w http.ResponseWriter, statusCode int, erro error) {

	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: erro.Error(),
	})

}
