package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// Client struct represents a single client
type Client struct {
	conn     net.Conn
	name     string
	messages chan string
}

// ChatServer struct represents the server that manages multiple clients
type ChatServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan string
	mu         sync.Mutex
	history    []string
}

var (
	port = "8989" // default port
	linuxLogo = `
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       \| \`' \\Zq
_)      \\.___.,|     .'
\\____   )MMMMMP|   .'
     \`-\'       '--'

)

func main() {
	// Set the port from arguments or use the default
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Initialize and start the server
	server := NewChatServer()
	go server.ListenAndServe()

	fmt.Printf("Listening on port :%s\n", port)
	select {}
}

// NewChatServer initializes the chat server
func NewChatServer() *ChatServer {
	return &ChatServer{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan string),
		history:    make([]string, 0),
	}
}

// ListenAndServe listens on the specified port and accepts incoming client connections
func (s *ChatServer) ListenAndServe() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	defer listener.Close()

	go s.handleMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err)
			continue
		}
		client := &Client{
			conn:     conn,
			messages: make(chan string),
		}
		go s.handleNewClient(client)
	}
}

// handleNewClient manages new client connections, prompts for name, and handles joining the chat
func (s *ChatServer) handleNewClient(client *Client) {
	client.conn.Write([]byte(linuxLogo + "\n"))
	client.conn.Write([]byte("[ENTER YOUR NAME]: "))

	reader := bufio.NewReader(client.conn)
	name, _ := reader.ReadString('\n')
	client.name = strings.TrimSpace(name)

	if client.name == "" {
		client.conn.Write([]byte("Name cannot be empty. Disconnecting...\n"))
		client.conn.Close()
		return
	}

	// Register the client
	s.register <- client

	// Send chat history to the new client
	for _, msg := range s.history {
		client.conn.Write([]byte(msg + "\n"))
	}

	go s.handleClientInput(client)
	go s.sendClientMessages(client)
}

// handleClientInput reads input from a client and broadcasts it
func (s *ChatServer) handleClientInput(client *Client) {
	reader := bufio.NewReader(client.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			s.unregister <- client
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, client.name, message)

		s.broadcast <- formattedMessage
	}
}

// sendClientMessages sends messages to a client
func (s *ChatServer) sendClientMessages(client *Client) {
	for message := range client.messages {
		client.conn.Write([]byte(message + "\n"))
	}
}

// handleMessages handles broadcasting messages to all clients and managing registration/unregistration
func (s *ChatServer) handleMessages() {
	for {
		select {
		case client := <-s.register:
			s.mu.Lock()
			if len(s.clients) >= 10 {
				client.conn.Write([]byte("Max clients reached. Disconnecting...\n"))
				client.conn.Close()
			} else {
				s.clients[client] = true
				joinMsg := fmt.Sprintf("[%s] has joined the chat...", client.name)
				s.broadcast <- joinMsg
			}
			s.mu.Unlock()

		case client := <-s.unregister:
			s.mu.Lock()
			if _, ok := s.clients[client]; ok {
				leaveMsg := fmt.Sprintf("[%s] has left the chat...", client.name)
				s.broadcast <- leaveMsg
				delete(s.clients, client)
				client.conn.Close()
			}
			s.mu.Unlock()

		case message := <-s.broadcast:
			// Add message to history
			s.history = append(s.history, message)

			// Broadcast message to all clients
			s.mu.Lock()
			for client := range s.clients {
				select {
				case client.messages <- message:
				default:
					close(client.messages)
					delete(s.clients, client)
				}
			}
			s.mu.Unlock()
		}
	}
}
