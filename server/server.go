package server

import (
	"encoding/json"
	"net/http"
)

type Weather struct {
	City     string `json:"city"`
	Forecast string `json:"forecast"`
}

func GetWeather(url string) (*Weather, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		panic(err)
	}

	return &weather, nil
}
