package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func MockloadAPIKey() (string, error) {
	_ = godotenv.Load()
	apiKey := os.Getenv("GITHUB_APIKEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key not found in environment variables")
	}

	return apiKey, nil
}
func TestGetRecentIssuesByLanguage(t *testing.T) {
	// Load environment variables from .env file
	apiKey, err := MockloadAPIKey()
	if err != nil {
		// Handle the error (e.g., log it and exit)
		fmt.Printf("Error loading API key: %v\n", err)
		os.Exit(1)
	}

	// Mock server to simulate GitHub API responses
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "token "+apiKey, r.Header.Get("Authorization"))

		repositories := []map[string]interface{}{
			{
				"name":              "repo1",
				"description":       "description1",
				"open_issues_count": 5,
				"html_url":          "https://github.com/user/repo1",
				"pushed_at":         time.Now().AddDate(0, -3, 0).Format(time.RFC3339),
				"stargazers_count":  10,
				"language":          "Go",
			},
			{
				"name":              "repo2",
				"description":       "description2",
				"open_issues_count": 2,
				"html_url":          "https://github.com/user/repo2",
				"pushed_at":         time.Now().AddDate(0, -1, 0).Format(time.RFC3339),
				"stargazers_count":  20,
				"language":          "Go",
			},
		}

		response := struct {
			Items []map[string]interface{} `json:"items"`
		}{Items: repositories}

		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	// Override the baseURL with the mock server URL

	t.Run("successful response", func(t *testing.T) {
		repos, err := GetRecentIssuesByLanguage("Go")
		assert.NoError(t, err)
		assert.Len(t, repos, 2)
		assert.Equal(t, "repo1", repos[0].Name)
		assert.Equal(t, "repo2", repos[1].Name)
	})

	t.Run("API key missing", func(t *testing.T) {
		os.Setenv("GITHUB_APIKEY", "")
		_, err := GetRecentIssuesByLanguage("Go")
		assert.Error(t, err)
		assert.Equal(t, "API key not found in environment variables", err.Error())
		os.Setenv("GITHUB_APIKEY", apiKey) // Reset the API key for other tests
	})
}
