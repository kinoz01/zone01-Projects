package server

import (
	"fmt"
	"log"
	"net"
	"os"
)

func HandleConnections(listener net.Listener) {
	
	// Remove cache file and server logs
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

	// Send previous messages to the new client
	cache, err := os.ReadFile(fmt.Sprintf("chat:%s.txt", Port))
	if err != nil {
		fmt.Fprint(conn, "I can't load chat history, this is due to an internal server error.\n")
		ServerLogs.WriteString(err.Error())
	}

	PrintLastMessages(cache, conn)

	// Listen for incoming messages
	for {

		// Send prompt after join message
		prompt := fmt.Sprintf("[%s][%s]: ", timeStamp(), name)
		fmt.Fprint(conn, prompt)

		if !scanner.Scan() {
			break // Exit if client disconnects
		}

		msg := scanner.Text()
		if msg == "/name" {
            // handle name change
            var err error
            name, err = ChangeClientName(conn, scanner, name)
            if err != nil {
				ServerLogs.WriteString(err.Error())
                break // Handle disconnect or other errors
            }
            continue
        }

		msg = MakePrintable(msg) 

		// Broadcast the message to all clients except the sender
		Broadcast <- Message{Sender: conn, Content: msg, Name: name}
	}

	// Handle client leaving
	Mu.Lock()
	delete(Clients, conn)
	Mu.Unlock()

	defer conn.Close()
	defer ServerLogs.WriteString(fmt.Sprintf("%s disconnected\n", name))

	Broadcast <- Message{Sender: conn, Content: fmt.Sprintf("\n%s has left the chat...", name), Name: name}
}
