package apiserver

import (
	"groupie/server"
	"log"
	"net/http"
)


func Serve() {
	var port server.Ports
	port.InitialisePorts()

	apimux := http.NewServeMux()
	apimux.HandleFunc("/groupie", DataHandler)

	if err := http.ListenAndServe(":"+port.ApiPort, apimux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
