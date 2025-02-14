package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	s "github.com/devinargenta/gobookshelf/structs"
)

type API struct {
	// API URL
	URL   string
	Token string
}

var Client = &http.Client{
	Timeout: 10 * time.Second,
}

func (api *API) Get(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", api.URL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+api.Token)
	res, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return body, nil
}

// getLibraries fetches the list of libraries from the API.
func (api *API) GetLibraries() ([]s.Library, error) {
	var libraries s.LibraryCollection
	body, err := api.Get("libraries")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &libraries); err != nil {
		return nil, err
	}
	return libraries.Libraries, nil
}

func (api *API) GetPersonalized(l string) (s.Result, error) {
	var result s.Result
	body, err := api.Get("libraries/" + l + "/personalized")
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(body, &result.Items); err != nil {
		return result, err
	}
	return result, nil
}
