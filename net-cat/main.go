package main

import (
	"TCPChat/server"
)

func main() {

	listener := server.StartServer()
	if listener == nil {
		return
	}
	server.ServeAndHandle(listener)
}
