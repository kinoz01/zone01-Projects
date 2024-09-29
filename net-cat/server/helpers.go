package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// Return a string representing the global format of a time.
func timeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Check if a string contains only printable characters
func IsPrintable(s string) bool {
	for _, r := range s {
		// ASCII printable characters range from 32 (space) to 126 (~)
		if r < 32 || r == 127 {
			return false
		}
	}
	return true
}

// Remove unprintable characters from the broadcasted message.
func MakePrintable(msg string) (result string) {
	for _, r := range msg {
		// ASCII printable characters range from 32 (space) to 126 (~)
		if r < 32 || r == 127 {
			continue
		}
		result += string(r)
	}
	return result
}

// Check if username already used.
func UsedName(s string) bool {
	for _, name := range Clients {
		if name == s {
			return true
		}
	}
	return false
}

// Write cache file to conn.
func PrintLastMessages(cache []byte, conn net.Conn) {
	_, err := conn.Write(cache)
	if err != nil {
		ServerLogs.WriteString(err.Error())
	}
}

// Change client name when it types "/name"
func ChangeClientName(conn net.Conn, scanner *bufio.Scanner, currentName string) (string, error) {
	// Prompt the client to enter a new name
	fmt.Fprint(conn, "Enter your new name: ")

	if !scanner.Scan() {
		return currentName, fmt.Errorf("client disconnected")
	}

	newName := scanner.Text()

	// Validate the new name
	if UsedName(newName) {
		conn.Write([]byte("Name is already taken. Try again with a different name.\n"))
		return currentName, nil
	}
	if !IsPrintable(newName) {
		conn.Write([]byte("Invalid name. Try again with a different name.\n"))
		return currentName, nil
	}

	// Lock the client map and update the client's name
	Mu.Lock()
	Clients[conn] = newName // Assign the new name
	Mu.Unlock()

	conn.Write([]byte(fmt.Sprintf("Your name has been changed to %s.\n", newName)))

	ServerLogs.WriteString(fmt.Sprintf("%s has changed his name to: %s.\n", currentName, newName))

	Broadcast <- Message{Sender: conn, Content: fmt.Sprintf("\n%s has changed his name to %s.", currentName, newName), Name: newName}

	return newName, nil
}

// Create cache and logs files
func CreateCacheAndLogs(Port string) {
	var err error
	CacheFile, err = os.Create(fmt.Sprintf("chat:%s.txt", Port))
	if err != nil {
		log.Fatal(err)
	}
	ServerLogs, err = os.Create(fmt.Sprintf("logs:%s.txt", Port))
	if err != nil {
		log.Fatal(err)
	}
	ServerLogs.WriteString(fmt.Sprintf("Server Started at: %s \n", Port))
}
