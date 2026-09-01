package main

import (
	"log"
)

func main() {
	apikey := ""
	model := "mistral-tiny"
	client := New(apikey, model)
	_, err := client.message("Hi, What is the date today!")
	if err != nil {
		log.Println("error:", err)
	}
}
