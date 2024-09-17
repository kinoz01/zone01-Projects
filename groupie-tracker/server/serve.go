package server

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Serve() {
	var port Ports
	port.InitialisePorts()

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.HandleFunc("/css/", FileHandler)
	http.HandleFunc("/js/", FileHandler)

	log.Printf("Starting server at http://127.0.0.1:%s", port.Port)
	if err := http.ListenAndServe(":"+port.Port, nil); err != nil {
		if strings.Contains(err.Error(), "address already in use") {
			p, _ := strconv.Atoi(port.Port)
			for {
				log.Printf("Address %d is already in use", p)
				p++
				log.Printf("Starting server at http://127.0.0.1:%d", p)
				http.ListenAndServe(":"+strconv.Itoa(p), nil)
				if p == 65000 {
					log.Fatalf("Server failed to start: there is no available ports in your machine")
					break
				}
			}
		}
		log.Fatalf("Server failed to start: %v", err)
	}
}
