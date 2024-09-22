package server

import (
	"fmt"
	"net"
	"time"
)

func timeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Check if a string contains only printable characters
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
	for _, name := range Clients {
		if name == s {
			return true
		}
	}
	return false
}

func PrintLastMessage(last []byte, conn net.Conn) {
	_, err := conn.Write(last)
	if err != nil {
		fmt.Println(err)
	}
}

func PrintClientsInfo(name string, conn net.Conn) {
	// Get the client's remote address
	clientAddr := conn.RemoteAddr().String()

	// Print the full address (IP:Port)
	fmt.Printf("Client %s connected from: %s\n", name, clientAddr)
}
