package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type API struct {
	// API URL
	URL   string
	Token string
}

func (api *API) get(path string) []byte {
	fmt.Println(api)
	req, err := http.NewRequest("GET", api.URL+path, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+api.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// why is api.Token not available here?
	// why is api.URL not available here?
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func (api *API) getPersonalized(l string) (Result, error) {
	var r Result
	body := api.get("libraries/" + l + "/personalized")
	if err := json.Unmarshal(body, &r.Entities); err != nil {
		return Result{}, err
	}
	return r, nil
}

// getLibraries fetches the list of libraries from the API.
func (api *API) getLibraries() ([]Library, error) {
	libs := Libraries{}
	body := api.get("libraries")
	if err := json.Unmarshal(body, &libs); err != nil {
		return nil, err
	}
	return libs.Libraries, nil
}
