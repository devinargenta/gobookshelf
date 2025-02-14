package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devinargenta/devinargenta/structs"
	"github.com/stretchr/testify/assert"
)

func TestAPI_get(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		statusCode int
		body       string
		wantErr    bool
	}{
		{
			name:       "successful request",
			path:       "/test",
			statusCode: http.StatusOK,
			body:       `{"key":"value"}`,
			wantErr:    false,
		},
		{
			name:       "request failed with status code",
			path:       "/test",
			statusCode: http.StatusInternalServerError,
			body:       `{"error":"internal server error"}`,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.body))
			}))
			defer server.Close()

			api := &API{
				URL:   server.URL,
				Token: "test-token",
			}

			got, err := api.Get(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && string(got) != tt.body {
				t.Errorf("API.get() = %v, want %v", string(got), tt.body)
			}
		})
	}
}
func TestGetLibraries(t *testing.T) {
	expectedLibraries := []structs.Library{
		{ID: "1"},
		{ID: "2"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/libraries", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		libraries := structs.LibraryCollection{Libraries: expectedLibraries}
		json.NewEncoder(w).Encode(libraries)
	}))
	defer server.Close()

	api := &API{
		URL:   server.URL + "/",
		Token: "test-token",
	}

	libraries, err := api.GetLibraries()
	assert.NoError(t, err)
	assert.Equal(t, expectedLibraries, libraries)
}

func TestAPI_GetPersonalized(t *testing.T) {
	tests := []struct {
		name       string
		libraryID  string
		statusCode int
		body       string
		wantErr    bool
	}{
		{
			name:       "successful request",
			libraryID:  "1",
			statusCode: http.StatusOK,
			body:       `[{"id":"1","name":"Item1"},{"id":"2","name":"Item2"}]`,
			wantErr:    false,
		},
		{
			name:       "request failed with status code",
			libraryID:  "1",
			statusCode: http.StatusInternalServerError,
			body:       `{"error":"internal server error"}`,
			wantErr:    true,
		},
		{
			name:       "invalid JSON response",
			libraryID:  "1",
			statusCode: http.StatusOK,
			body:       `invalid json`,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.body))
			}))
			defer server.Close()

			api := &API{
				URL:   server.URL + "/",
				Token: "test-token",
			}

			got, err := api.GetPersonalized(tt.libraryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.getPersonalized() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				var expectedResult structs.Result
				if err := json.Unmarshal([]byte(tt.body), &expectedResult.Items); err != nil {
					t.Fatalf("failed to unmarshal expected body: %v", err)
				}
				if len(got.Items) != len(expectedResult.Items) {
					t.Errorf("API.getPersonalized() = %v, want %v", got.Items, expectedResult.Items)
				}
			}
		})
	}
}
