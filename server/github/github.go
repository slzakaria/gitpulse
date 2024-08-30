package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const baseURL = "https://api.github.com"

type Repository struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	OpenIssues   int       `json:"open_issues_count"`
	URL          string    `json:"html_url"`
	LastActivity time.Time `json:"-"`
	LastPushedAt string    `json:"pushed_at"`
	Stars        int       `json:"stargazers_count"`
	Language     string    `json:"language"`
}

func loadAPIKey() (string, error) {
	_ = godotenv.Load()
	apiKey := os.Getenv("GITHUB_APIKEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key not found in environment variables")
	}

	return apiKey, nil
}

func GetRecentIssuesByLanguage(language string) ([]Repository, error) {
	apiKey, err := loadAPIKey()
	if err != nil {
		fmt.Printf("Error loading API key: %v\n", err)
		os.Exit(1)
	}

	sixMonthsAgo := time.Now().AddDate(0, -6, 0).Format("2006-01-02T15:04:05Z")

	apiUrl := fmt.Sprintf("%s/search/repositories?q=language:%s+is:public+pushed:>%s+open:issues&sort=updated&order=desc", baseURL, language, sixMonthsAgo)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set GitHub API token
	req.Header.Set("Authorization", apiKey)

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result struct {
		Items []map[string]interface{} `json:"items"`
	}

	if error := json.NewDecoder(response.Body).Decode(&result); error != nil {
		return nil, error
	}

	// Parse and format the "pushed_at" field to get the yyyy-mm-dd format
	var repositories []Repository
	for _, item := range result.Items {
		pushedAt, ok := item["pushed_at"].(string)
		if !ok {
			continue
		}

		pushedTime, err := time.Parse(time.RFC3339, pushedAt)
		if err != nil {
			continue
		}

		repository := Repository{
			Name:         item["name"].(string),
			Description:  item["description"].(string),
			OpenIssues:   int(item["open_issues_count"].(float64)),
			URL:          item["html_url"].(string),
			LastActivity: pushedTime,
			LastPushedAt: pushedTime.Format("2006-01-02"),
			Stars:        int(item["stargazers_count"].(float64)),
			Language:     item["language"].(string),
		}

		repositories = append(repositories, repository)
	}

	// Filter repositories with open issues greater than 1
	filteredResult := []Repository{}
	for _, repo := range repositories {
		if repo.OpenIssues > 1 {
			filteredResult = append(filteredResult, repo)
		}
	}

	return filteredResult, nil
}
