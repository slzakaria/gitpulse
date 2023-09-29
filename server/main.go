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
		log.Fatalf("Error loading .env file inside main.go: %v", err)
	}

	PORT := os.Getenv("PORT")
	apik := os.Getenv("GITHUB_API_KEY")
	fmt.Println(apik, PORT)

	// Set up API routes
	router := api.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "https://gittrackr.vercel.app", "https://gitpulse-zackaria-sl.vercel.app"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Use the CORS middleware
	handler := c.Handler(router)

	port := ":3000"

	log.Printf("Server is running on port %s ...\nFetching github issues opened in the last 6 months", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
