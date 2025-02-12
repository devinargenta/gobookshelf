package main

import (
	"encoding/json"
)

// type Libraries struct {
// 	Libraries []Library `json:"libraries"`
// }

// type Library struct {
// 	ID    string `json:"id"`
// 	Name  string `json:"name"`
// 	Items []Item `json:"items"`
// }

type Stats struct {
	TotalTime json.Number     `json:"totalTime"`
	Today     json.Number     `json:"today"`
	Items     map[string]Item `json:"items"`
}
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Entities struct {
	ID string `json:"id"`
}
type MediaMetadata struct {
	AuthorName string `json:"authorname"`
	Title      string `json:"title"`
}
type Media struct {
	ID       string        `json:"id"`
	Metadata MediaMetadata `json:"metadata"`
}

type Result struct {
	Entities []Items `json:"entities"`
}
type Items struct {
	ID       string   `json:"id"`
	Label    string   `json:"label"`
	Entities []Entity `json:"entities"`
}

type Entity struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	MediaType string `json:"mediaType"`
	Metadata  Media  `json:"media"`
}

type Library struct {
	ID string `json:"id"`
}

type Libraries struct {
	Libraries []Library `json:"libraries"`
}
