package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"codetrackr/api"
)

func main() {
	fmt.Println("Started the go server successfully")
	router := api.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://gittrackr.vercel.app", "https://gitpulse-zackaria-sl.vercel.app", "https://git-tracker-front.onrender.com", "https://git-tracker.onrender.com"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"*"},
	})

	// Use the CORS middleware
	handler := c.Handler(router)

	port := ":3000"

	log.Printf("Server is running  .....\nFetching github issues opened in the last 6 months")
	log.Fatal(http.ListenAndServe(port, handler))
}
