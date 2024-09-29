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
	// sigChan <- os.Interrupt // We can use this even if the syntax is valid, a signal is special case that we have Notify function to deal with it.

	go func() {
		// Send a value to a channel: channel <- value
		// Receive a value from a channel: var := <-channel
		v := <-sigChan
		fmt.Printf("\nReceived %v signal, cleaning up...\n", v)

		// Close and delete cache file
		if CacheFile != nil  {
			CacheFile.Close()
			err := os.Remove(CacheFile.Name())
			if err != nil {
				log.Fatalf("Error deleting cache file: %v\n", err)
			} else {
				fmt.Println("Cache file deleted successfully.")
			}
		}
		// Close and delete server logs
		if ServerLogs != nil  {
			ServerLogs.Close()
			err := os.Remove(ServerLogs.Name())
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
