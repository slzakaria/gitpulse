package main

import (
	"codetrackr/api"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Started the go server successfully")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	// Set up API routes
	router := api.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Use the CORS middleware 
	handler := c.Handler(router)

	port := PORT
	log.Printf("Server is running on port %s ...\nFetching github issues opened in the last 3 months", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
