package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	// Define API routes here
	router.Get("/api/issues/{language}", GetRecentIssuesByLanguageHandler)

	return router
}
