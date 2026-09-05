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
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		panic("Scanner break")
	}

	for false {
		fmt.Print("user: ")
		if !scanner.Scan() {
			break
		}

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

		fmt.Println("assistant:", res)
	}
}

func readFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
