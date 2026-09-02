package main

import (
	"fmt"
	"log"
)

// "fmt"
// "log"
// "os"

// "github.com/joho/godotenv"

func main() {
	city := "london"
	data, err := WeatherOfACity(city)
	if err != nil {
		log.Fatal("city weather fail:", err)
	}

	fmt.Println(city, ":", data)

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("error load .env file")
	// }
	// apikey := os.Getenv("APIKEY")
	// model := "mistral-tiny"
	// client := New(apikey, model)
	// res, err := client.message("Hi, What is the date today!")
	// if err != nil {
	// 	log.Println("error:", err)
	// }

	// fmt.Println("res:", res)

}
