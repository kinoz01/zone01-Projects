package server

import (
	"fmt"
	"net"
	"os"
)

// HandleConnections accepts incoming connections from clients
// and spawns a goroutine to handle each client connection.
func HandleConnections(listener net.Listener) {
	// Remove cache file and server logs
	RemoveCahe()
	go BroadcastMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			ServerLogs.WriteString(err.Error() + "\n")
			continue
		}

		// Handle each connected client in its own goroutine
		go HandleClients(conn)
	}
}

// HandleClients manages the interaction with a single client.
// It reads messages from the client and broadcasts them to all other clients.
func HandleClients(conn net.Conn) {
	name, scanner := AcceptNewClient(conn)
	if scanner == nil {
		return
	}

	// Send previous messages to the new client
	cache, err := os.ReadFile(CacheFile.Name())
	if err != nil {
		fmt.Fprint(conn, "I can't load chat history, this is due to an internal server error.\n")
		ServerLogs.WriteString(err.Error() + "\n")
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
				ServerLogs.WriteString(err.Error() + "\n")
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
