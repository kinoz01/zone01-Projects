package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Check port number and start server
func StartServer() net.Listener{
	var port string

	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return nil
	}
	if port == "" {
		port = "8989"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port", port)
	return listener
}
