package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/devinargenta/devinargenta/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	api := &api.API{
		URL:   os.Getenv("API_ROOT"),
		Token: os.Getenv("TOKEN"),
	}

	libs, err := api.GetLibraries()
	if err != nil {
		log.Fatal(err)
	}
	// get the first item
	if len(libs) == 0 {
		log.Println("No libraries found")
		return
	}

	l, err := api.GetPersonalized(libs[0].ID)
	if err != nil {
		log.Printf("Error getting personalized data for library ID %s: %v", libs[0].ID, err)
		return
	}

	printJSON(l.Items)
}

// prints json nicely for me
func printJSON(s any) {
	b, err := json.MarshalIndent(s, ">", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
