package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	City string `json:"city"`
	Temp string `json:"temp"`
}

func WeatherOfACity(city string) (Data, error) {
	url := fmt.Sprintf("http://localhost:8080/api/v1/weather?city=%s", city)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("http request error")
		return Data{}, err
	}
	defer resp.Body.Close()

	var data Data
	if err2 := json.NewDecoder(resp.Body).Decode(&data); err2 != nil {
		log.Fatal("request body decode error")
		return Data{}, err2
	}
	return data, nil
}
