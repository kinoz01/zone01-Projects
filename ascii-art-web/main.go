package main

import (
	"asciiArt/server"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/ascii-art", server.AsciiArtHandler)

	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
