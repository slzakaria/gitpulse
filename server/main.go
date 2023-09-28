package main

import (
	"codetrackr/api"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Set up API routes
	router := api.SetupRoutes()

	// Create a new CORS middleware instance
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Replace with your frontend URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Use the CORS middleware with your router
	handler := c.Handler(router)

	port := ":8080"
	log.Printf("Server is running on port %s ...\nFetching github issues opened in the last 3 months", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
