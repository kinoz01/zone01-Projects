package asciiart

import (
	"fmt"
	"regexp"
	"strings"
)

func UerArgs(args []string) (userText, font, outputFile string) {
	reOutput := regexp.MustCompile(`\A--output=(\S+.txt)$`)

	switch len(args) {
	case 1:
		if strings.HasPrefix(args[0], "--output") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return userText, font, outputFile
		}
		userText = args[0]
		font = "standard"
	case 2:
		if reOutput.MatchString(args[0]) {
			userText = args[1]
			font = "standard"
			outputFile = reOutput.FindStringSubmatch(args[0])[1]
		} else {
			userText = args[0]
			font = args[1]
		}
	case 3:
		if reOutput.MatchString(args[0]) {
			userText = args[1]
			font = args[2]
			outputFile = reOutput.FindStringSubmatch(args[0])[1]			
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		}
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}
	return userText, font, outputFile
}
