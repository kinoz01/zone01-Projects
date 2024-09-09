package server

import (
	"log"
	"net/http"
)

func Serve() {
	var port Ports
	port.InitialisePorts()

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.HandleFunc("/css/", CSSHandler)
	http.HandleFunc("/js/", JSHandler)

	log.Printf("Starting server at http://127.0.0.1:%s", port.Port)
	if err := http.ListenAndServe(":"+port.Port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
