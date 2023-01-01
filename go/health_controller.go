package main

import (
	"net/http"
	"os"
)

func getHealthEndpoint(w http.ResponseWriter, req *http.Request) {
	name, err := os.Hostname()
	status := true
	if err != nil {
		status = false
	}
	health := Health{Status: status, Host: name}
	ResponseJson(w, health.status(), health)
}

type Health struct {
	Status bool   `json:"status"`
	Host   string `json:"hostname"`
}

func (health Health) status() int {
	if health.Status != true {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
