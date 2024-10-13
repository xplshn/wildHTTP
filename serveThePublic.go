package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	ip, err := getPublicIP()
	if err != nil {
		log.Fatalf("Failed to fetch public IP: %v", err)
	}

	fmt.Printf("Server is running on public IP: %s, port: %s\n", ip, port)

	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Printf("Serving files at http://%s:%s\n", ip, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
