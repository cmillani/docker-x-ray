package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Weather struct {
	TemperatureC float64 `json:"temperatureC"`
	WindSpeedKmH float64 `json:"windSpeedKmH"`
}

func didReceiveWeather(weather Weather) error {
	fmt.Printf("Ingesting %v\n", weather)
	return nil
}

func retrieveData() ([]Weather, error) {
	resp, err := http.Get("https://622e1cd18d943bae348ebf0c.mockapi.io/api/weather")
	if err != nil {
		return nil, err
	}

	var weatherList []Weather
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&weatherList)
	if err != nil {
		return nil, err
	}

	return weatherList, nil
}
