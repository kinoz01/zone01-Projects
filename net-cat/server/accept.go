package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func AcceptNewClient(conn net.Conn) (string, *bufio.Scanner) {

	Mu.Lock()
	if len(Clients) >= MaxClients {
		conn.Write([]byte("Chat room is full. Please try again later.\n"))
		Mu.Unlock()
		conn.Close() // Close connection if chat room is full
		return "", nil
	}
	Mu.Unlock()

	file, err := os.ReadFile("bitri9.txt")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write(file)

	conn.Write([]byte("Welcome to TCP-Chat!\n[ENTER YOUR NAME]: "))
	name := ""
	scanner := bufio.NewScanner(conn)

	for {
		if !scanner.Scan() {
			return "", nil // Client disconnected
		}

		name = scanner.Text()
		if !IsPrintable(name) {
			conn.Write([]byte("[Please Enter a valid name]: "))
			continue
		}

		Mu.Lock()
		if UsedName(name) {
			conn.Write([]byte("[Name already used, please enter a new name]: "))
			Mu.Unlock()
			continue
		}

		// Valid name, proceed
		Clients[conn] = name
		Mu.Unlock()
		break
	}

	Broadcast <- Message{Sender: conn, Content: fmt.Sprintf("\n%s has joined the chat...", name), Name: name}

	return name, scanner
}
