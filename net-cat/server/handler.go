package server

import (
	"fmt"
	"log"
	"net"
	"os"
)

func ServeAndHandle(listener net.Listener) {
	
	// Remove cache file (server logs)
	RemoveCahe()
	go BroadcastMessages()	
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle each connection in a separate goroutine
		go HandleClients(conn)
	}
}


func HandleClients(conn net.Conn) {

	name, scanner := AcceptNewClient(conn)
	if scanner == nil {
		return
	}

	PrintClientsInfo(name, conn)

	// Send previous messages to the new client
	logs, err := os.ReadFile(fmt.Sprintf("chat:%s.txt", Port))
	if err != nil {
		fmt.Fprint(conn, "I can't find previous messsages, this is due to an internal server error.\n")
	}

	PrintLastMessage(logs, conn)

	// Listen for incoming messages
	for {

		// Send prompt after join message
		prompt := fmt.Sprintf("[%s][%s]: ", timeStamp(), name)
		fmt.Fprint(conn, prompt)

		if !scanner.Scan() {
			break // Exit if client disconnects
		}

		msg := scanner.Text()
		if !IsPrintable(msg) {
			conn.Write([]byte("[Please Enter a valid message]: \n"))
			continue
		}

		// Broadcast the message to all clients except the sender
		Broadcast <- Message{Sender: conn, Content: msg, Name: name}
	}

	// Handle client leaving
	Mu.Lock()
	delete(Clients, conn)
	Mu.Unlock()

	defer conn.Close()
	defer fmt.Printf("%s disconnected\n", name)

	Broadcast <- Message{Sender: conn, Content: fmt.Sprintf("\n%s has left the chat...", name), Name: name}
}
