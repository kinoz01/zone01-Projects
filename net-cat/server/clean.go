package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

// RemoveCahe sets up a handler to catch OS signals (like Ctrl+C)
// and clean cachefile created by the conversation.
func RemoveCahe() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal, cleaning up...")

		// Close and delete cache file
		if CacheFile != nil {
			CacheFile.Close()
			err := os.Remove(CacheFile.Name())
			if err != nil {
				log.Fatalf("Error deleting server logs: %v\n", err)
			} else {
				fmt.Println("Server logs deleted successfully.")
			}
		}

		// Exit the program gracefully
		os.Exit(0)
	}()
}
