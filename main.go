package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// city := "london"
	// data, err := WeatherOfACity(city)
	// if err != nil {
	// 	log.Fatal("city weather fail:", err)
	// }

	// fmt.Println(city, ":", data)
	for false {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("user: ")
		scanner.Scan()
		prompt := scanner.Text()

		if strings.ToLower(prompt) == "quit" || strings.ToLower(prompt) == "exit" {
			break
		}

		if err := godotenv.Load(); err != nil {
			log.Fatal("error load .env file")
		}
		apikey := os.Getenv("APIKEY")
		model := "mistral-tiny"
		client := New(apikey, model)
		res, err := client.message(prompt)
		if err != nil {
			log.Println("error:", err)
		}

		fmt.Println("assistant:", res.Choices[0].Message.Content)
	}
}

func readFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	return string(bytes), err
}
