package server

import (
	"log"
	"net"
	"net/http"
)

var Port Ports
var APILinks *ApiLinks

func Serve() {
	Port.InitialisePorts()
	InitialiseApiLinks()

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.HandleFunc("/css/", FileHandler)
	http.HandleFunc("/js/", FileHandler)

	listener, err := net.Listen("tcp", ":"+Port.Port)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
		return
	}
	_, p, _ := net.SplitHostPort(listener.Addr().String())

	log.Printf("Starting server at http://127.0.0.1:%s", p)

	if err = http.Serve(listener, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
