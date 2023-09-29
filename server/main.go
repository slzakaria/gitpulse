package main

import (
	"codetrackr/api"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	fmt.Println("Started the go server successfully")

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	// PORT := os.Getenv("PORT")

	// Set up API routes
	router := api.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "https://gittrackr.vercel.app/"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Use the CORS middleware
	handler := c.Handler(router)

	port := ":3000"
	fmt.Println("Port is :", port)

	log.Printf("Server is running on port %s ...\nFetching github issues opened in the last 3 months", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
