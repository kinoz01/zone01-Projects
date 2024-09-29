package server

import "fmt"

// Broadcast handler that sends messages to all clients except the sender
func BroadcastMessages() {
	for {
		message := <-Broadcast
		
		formattedMessage := fmt.Sprintf("[%s][%s]: %s\n", timeStamp(), message.Name, message.Content)
		CacheFile.WriteString(formattedMessage)

		Mu.Lock()
		for conn, username := range Clients {
			if conn != message.Sender {
				_, err := conn.Write([]byte("\n" + formattedMessage))
				if err != nil {
					ServerLogs.WriteString(fmt.Sprintf("Error writing to connection: %v\n", err))
					conn.Close()
					delete(Clients, conn)
					continue
				}
				prompt := fmt.Sprintf("[%s][%s]: ", timeStamp(), username)
				fmt.Fprint(conn, prompt)
			}
		}
		Mu.Unlock()
	}
}
