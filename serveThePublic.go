package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define the port for the server, fallback to port 8080
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	// Serve files from the ./public directory
	http.Handle("/", http.FileServer(http.Dir("./public")))

	// Log the server starting
	log.Printf("Server starting on port %s\n", port)

	// Start serving
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
