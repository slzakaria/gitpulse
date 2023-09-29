package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func GetRecentIssuesByLanguage(language string) ([]Repository, error) {
	// Load environment variables from .env file
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// apiKey := os.Getenv("GITHUB_API_KEY")

	// Calculate the date three months ago from now and fetch repos with open issues during the timefrime
	threeMonthsAgo := time.Now().AddDate(0, -3, 0).Format("2006-01-02T15:04:05Z")

	apiUrl := fmt.Sprintf("%s/search/repositories?q=language:%s+is:public+pushed:>%s+open:issues&sort=updated&order=desc", baseURL, language, threeMonthsAgo)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set GitHub API token
	req.Header.Set("Authorization", "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")

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
