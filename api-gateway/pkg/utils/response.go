package utils

import (
	"encoding/json"
	"net/http"
)


type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{Error: err.Error()}
	jsonResponse, _ := json.Marshal(errorResponse) 

	w.Write(jsonResponse)
}


func SendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	responseBytes, err := json.Marshal(data)
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write(responseBytes)
}