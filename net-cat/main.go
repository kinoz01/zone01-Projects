package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var clients = make(map[net.Conn]string)
var messages = make([]string, 0)
var mu sync.Mutex
var broadcast = make(chan Message)

const maxClients = 2

// Message struct to store message content and the sender
type Message struct {
	Sender  net.Conn
	Content string
	Name    string
}

func main() {
	listener := StartServer()
	if listener == nil {
		return
	}
	defer listener.Close()

	// Start the broadcast handler in a goroutine

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle each connection in a separate goroutine
		go HandleClient(conn)
		go Broadcast()
	}

}

// Broadcast handler that sends messages to all clients except the sender

func Broadcast() {
	for {
		ms := <-broadcast
	
		formattedMessage := fmt.Sprintf("[%s][%s]: %s-----------\n", timeStamp(), ms.Name, ms.Content)
		mu.Lock()
		for conn, username := range clients {
			if conn != ms.Sender {
				_, err := conn.Write([]byte("\n" + formattedMessage))
				if err != nil {
					fmt.Printf("Error writing to connection: %v\n", err)
					conn.Close()
					delete(clients, conn)
				}			
				prompt := fmt.Sprintf("[%s][%s]: ", timeStamp(), username)
				fmt.Fprint(conn, prompt)
			}
		}
		mu.Unlock()
	}
}

func HandleClient(conn net.Conn) {
	mu.Lock()
	if len(clients) >= maxClients {
		conn.Write([]byte("Chat room is full. Please try again later.\n"))
		mu.Unlock()
		conn.Close() // Close connection if chat room is full
		return
	}
	mu.Unlock()

	defer conn.Close()

	conn.Write([]byte("Welcome to TCP-Chat!\n[ENTER YOUR NAME]: "))
	name := ""
	scanner := bufio.NewScanner(conn)

	for {
		if !scanner.Scan() {
			return // Client disconnected
		}

		name = scanner.Text()
		if !IsPrintable(name) {
			conn.Write([]byte("[Please Enter a valid name]: "))
			continue
		}

		mu.Lock()
		if UsedName(name) {
			conn.Write([]byte("[Name already used, please enter a new name]: "))
			mu.Unlock()
			continue
		}

		// Valid name, proceed
		clients[conn] = name
		mu.Unlock()
		break
	}

	broadcast <- Message{Sender: conn, Content: fmt.Sprintf("%s has joined the chat...", name), Name: name}

	// Send previous messages to the new client
	mu.Lock()
	for _, msg := range messages {
		conn.Write([]byte(msg + "\n"))
	}
	mu.Unlock()

	// Notify user of ready state for writing messages
	//conn.Write([]byte(fmt.Sprintf("%s: ", name))) // First prompt

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
			conn.Write([]byte("[Please Enter a valid message]: "))
			continue
		}

		// Format the message to show <name>: <message>
		//formattedMessage := fmt.Sprintf("[%s] %s: %s", timeStamp(), name, msg)

		// Append the message to history
		mu.Lock()
		messages = append(messages, msg)
		mu.Unlock()

		// Broadcast the message to all clients except the sender
		broadcast <- Message{Sender: conn, Content: msg, Name:  name}

		// Restore the user prompt after broadcasting the message
		//conn.Write([]byte(fmt.Sprintf("\n%s: ", name))) // Writing cursor restored after message is sent
	}

	// Handle client leaving
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	broadcast <- Message{Sender: conn, Content: fmt.Sprintf("%s has left the chat...", name), Name: name}
}

func timeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Function to check if a string contains only printable characters
func IsPrintable(s string) bool {
	for _, r := range s {
		// ASCII printable characters range from 32 (space) to 126 (~)
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

func UsedName(s string) bool {
	for _, name := range clients {
		if name == s {
			return true
		}
	}
	return false
}
