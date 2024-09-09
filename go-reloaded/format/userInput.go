package format

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Get user input after printing a message containging the flag match we are in (used to parse wrong flags).
// I removed most of wrong flags handling (user reactions) and only left the one with wrong signs in this version.
func GetUserInput(match string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("🚩 A flag pattern < %s > was found. Do you want to format it to a valid flag? (y/n): ", match)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "y" || input == "n" {
			return input
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
}

// Get user input after printing a message (used for cleaning spaces).
func GetUserInputPrompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "y" || input == "n" {
			return input
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
}
