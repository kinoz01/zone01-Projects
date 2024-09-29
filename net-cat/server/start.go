package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var (
	Clients    = make(map[net.Conn]string)
	Mu         sync.Mutex
	Broadcast  = make(chan Message)
	Quit       = make(chan bool)
	CacheFile  *os.File
	ServerLogs *os.File
	Port       string
)

const MaxClients = 10

// Message struct to store message content and the sender
type Message struct {
	Sender  net.Conn
	Content string
	Name    string
}

// Check port number and start server
func StartServer() net.Listener {
	// fmt.Print("\033[H\033[2J")

	if len(os.Args) == 2 {
		Port = os.Args[1]
	}
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $Port")
		return nil
	}
	if Port == "" {
		Port = "8989"
	}

	listener, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on Port", Port)

	CreateCacheAndLogs(Port)

	return listener
}
