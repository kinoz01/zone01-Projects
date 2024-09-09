package server

import (
	"log"
	"net/http"
	"os"
)

// Environment variables
var Port string
var StyleEnv string


// Set Handler and Start the server using defaultServeMux.
func NewServer() {
	StyleEnv = os.Getenv("STYLE")
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8088" // Default port
	}

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)
	http.HandleFunc("/css/", CSSHandler)
	log.Printf("Starting server on http://127.0.0.1:%s", Port)

	if err := http.ListenAndServe(":"+Port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
