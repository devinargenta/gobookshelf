package structs

import (
	"encoding/json"
)

type Stats struct {
	TotalTime json.Number     `json:"totalTime"`
	Today     json.Number     `json:"today"`
	Items     map[string]Item `json:"items"`
}
type Item struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	MediaType string `json:"mediaType"`
	Metadata  Media  `json:"media"`
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
	Items []EntityGroup `json:"items"`
}

type EntityGroup struct {
	ID       string   `json:"id"`
	Label    string   `json:"label"`
	Entities []Entity `json:"entities"`
	Type     string   `json:"type"`
}

type Entity struct {
	ID        string `json:"id"`
	MediaType string `json:"mediaType"`
	Metadata  Media  `json:"media"`
}

type ItemLabel struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type LibraryCollection struct {
	Libraries []Library `json:"libraries"`
}

type Library struct {
	ID string `json:"id"`
}
