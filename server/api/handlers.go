package api

import (
	"codetrackr/github"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// handles the GET /api/issues/{language} route
func GetRecentIssuesByLanguageHandler(w http.ResponseWriter, r *http.Request) {
	language := chi.URLParam(r, "language")

	issues, err := github.GetRecentIssuesByLanguage(language)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(issues); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
