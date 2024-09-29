package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// AcceptNewClient handles the connection of a new client to the chat server.
func AcceptNewClient(conn net.Conn) (string, *bufio.Scanner) {

	Mu.Lock()
	Clients[conn] = "temporaryName" // we add the client with a dummy name to check for the max condition.
	if len(Clients) > MaxClients {
		conn.Write([]byte("Chat room is full. Please try again later.\n"))
		Mu.Unlock()
		ServerLogs.WriteString(fmt.Sprintf("Chat room is full, client connection refused from: %s\n", conn.RemoteAddr().String()))
		conn.Close() // Close connection if chat room is full
		return "", nil
	}
	Mu.Unlock()


	file, err := os.ReadFile("bitri9.txt")
	if err != nil {
		ServerLogs.WriteString(err.Error() + "\n")
	}
	if _, err = conn.Write(file); err != nil{
		ServerLogs.WriteString(err.Error() + "\n")
	}

	conn.Write([]byte("Welcome to TCP-Chat!\n[ENTER YOUR NAME]: "))
	name := ""
	scanner := bufio.NewScanner(conn)

	for {
		if !scanner.Scan() {
			ServerLogs.WriteString(fmt.Sprintf("Client %s disconnected while choosing his name.\n", conn.RemoteAddr().String()))
			return "", nil // Client disconnected
		}

		name = scanner.Text()

		if !IsPrintable(name) {
			conn.Write([]byte("Please Enter a valid name: "))
			continue
		}

		Mu.Lock()
		if UsedName(name) {
			conn.Write([]byte("Name already used, please enter a new name: "))
			Mu.Unlock()
			continue
		}

		// Valid name, proceed
		Clients[conn] = name
		Mu.Unlock()
		break
	}

	// Log Client info.
	ServerLogs.WriteString(fmt.Sprintf("Client %s connected from: %s\n", name, conn.RemoteAddr().String()))

	Broadcast <- Message{Sender: conn, Content: fmt.Sprintf("\n%s has joined the chat...", name), Name: name}

	return name, scanner
}
