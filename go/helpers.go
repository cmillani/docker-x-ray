package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func ResponseError(w http.ResponseWriter, e error) {
	fmt.Printf("Error: %v", e)
	ResponseJson(w, http.StatusInternalServerError, nil)
}
