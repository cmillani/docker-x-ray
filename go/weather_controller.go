package main

import (
	"encoding/json"
	"net/http"
)

func weatherEndpoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		retrieveWeather(w, req)
	case http.MethodPost:
		createWeather(w, req)
	default:
		ResponseJson(w, http.StatusMethodNotAllowed, nil)
	}
}

func retrieveWeather(w http.ResponseWriter, req *http.Request) {
	resp, err := retrieveData()
	if err != nil {
		ResponseError(w, err)
		return
	}

	ResponseJson(w, http.StatusOK, resp)
}

func createWeather(w http.ResponseWriter, req *http.Request) {
	var weather Weather
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&weather)
	if err == nil {
		err = didReceiveWeather(weather)
	}
	if err != nil {
		ResponseError(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
