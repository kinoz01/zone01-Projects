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
var broadcast = make(chan string)

func main() {
	port := ":8989"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Listening on port", port)

	// Start the broadcast handler in a goroutine
	go handleBroadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle each connection in a separate goroutine
		go handleClient(conn)
	}
}

// Broadcast handler that sends messages to all clients
func handleBroadcast() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			client.Write([]byte(msg + "\n"))
		}
		mu.Unlock()
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Welcome to TCP-Chat!\n[ENTER YOUR NAME]: "))
	name := ""
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		name = scanner.Text()
		if name != "" {
			break
		}
		conn.Write([]byte("[ENTER YOUR NAME]: "))
	}

	mu.Lock()
	clients[conn] = name
	mu.Unlock()

	broadcast <- fmt.Sprintf("[%s] %s has joined the chat...", timeStamp(), name)

	// Send previous messages to the new client
	mu.Lock()
	for _, msg := range messages {
		conn.Write([]byte(msg + "\n"))
	}
	mu.Unlock()

	// Listen for incoming messages
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "" {
			continue
		}
		broadcastMsg := fmt.Sprintf("[%s][%s]: %s", timeStamp(), name, msg)
		mu.Lock()
		messages = append(messages, broadcastMsg)
		mu.Unlock()
		broadcast <- broadcastMsg
	}

	// Handle client leaving
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	broadcast <- fmt.Sprintf("[%s] %s has left the chat...", timeStamp(), name)
}

func timeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
