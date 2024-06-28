package main

import (
	"asciiArt/api"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/ascii-art", api.AsciiArtHandler)

	log.Println("Starting server on http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
