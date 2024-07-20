package main

import (
	"asciiArt/api"
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux() // Create a ServeMux

    srv := &http.Server{
        Addr:         ":8088",
        Handler:      mux,       // Attach the ServeMux to the server
        ReadTimeout:  3 * time.Second,
        WriteTimeout: 3 * time.Second,
        IdleTimeout:  3 * time.Second,
    }

    // Register handlers with the ServeMux
    mux.HandleFunc("/", api.HomeHandler)
    mux.HandleFunc("/ascii-art", api.AsciiArtHandler)

    log.Println("Starting server on http://127.0.0.1:8088")
    err := srv.ListenAndServe()
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
