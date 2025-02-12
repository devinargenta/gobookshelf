package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api := &API{
		URL:   os.Getenv("API_ROOT"),
		Token: os.Getenv("TOKEN"),
	}

	libs, err := api.getLibraries()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan []Items, len(libs))
	// get the frst item

	l, err := api.getPersonalized(libs[0].ID)
	if err != nil {
		log.Println(err)
	}
	results <- l.Entities

	go func() {
		close(results)
	}()

	for result := range results {
		for _, item := range result {
			printJSON(item)
		}
	}
}

// prints json nicely for me
func printJSON(s any) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
