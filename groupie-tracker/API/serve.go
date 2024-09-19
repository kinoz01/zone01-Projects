package apiserver

import (
	"groupie/server"
	"log"
	"net"
	"net/http"
)

// run and serve the api server (that will serve json data)
func Serve() {
	server.Port.InitialisePorts()

	apimux := http.NewServeMux()
	apimux.HandleFunc("/groupie", DataHandler)

	listener, err := net.Listen("tcp", ":"+server.Port.ApiPort)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
		return
	}

	_, p, _ := net.SplitHostPort(listener.Addr().String())
	server.Port.ApiPort = p

	log.Printf("Starting API server at http://127.0.0.1:%s", p)
	if err = http.Serve(listener, apimux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
