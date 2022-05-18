package main

import (
	"net/http"
)

func getHealthEndpoint(w http.ResponseWriter, req *http.Request) {
	health := Health{Status: true}
	ResponseJson(w, health.status(), health)
}

type Health struct {
	Status bool `json:"status"`
}

func (health Health) status() int {
	return http.StatusOK
}
