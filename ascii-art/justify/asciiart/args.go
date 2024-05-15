package asciiart

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	alignErr = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard"
)

func UserArgs(args []string) (userText, font, outputFile, alignement string, quit bool) {
	alignement = "left"
	font = "standard"
	reOutput := regexp.MustCompile(`\A--output=(\S+.txt)$`)
	reAlign := regexp.MustCompile(`\A--align=(center|left|right|justify)$`)

	switch len(args) {
	case 1:
		if strings.HasPrefix(args[0], "--output") {
			fmt.Println(alignErr)
			quit = true
		} else if strings.HasPrefix(args[0], "--align") {
			fmt.Println(alignErr)
		} else {
			userText = args[0]
		}
	case 2:
		if (reOutput.MatchString(args[0]) && reAlign.MatchString(args[1])) || (reOutput.MatchString(args[1]) && reAlign.MatchString(args[0])){
			fmt.Println(alignErr)
			quit = true
		} else if reOutput.MatchString(args[0]) && !strings.HasPrefix(args[1], "--align") {
			outputFile = reOutput.FindStringSubmatch(args[0])[1]
			userText = args[1]
		} else if reAlign.MatchString(args[0]) && !strings.HasPrefix(args[1], "--output") {
			alignement = reAlign.FindStringSubmatch(args[0])[1]
			userText = args[1]
		} else {
			fmt.Println("hey")
			userText = args[0]
			font = args[1]
		}
	case 3:
		if !reOutput.MatchString(args[0]) && !reAlign.MatchString(args[0]){
			fmt.Println(alignErr)
			quit = true
		} else if reOutput.MatchString(args[0]) && reAlign.MatchString(args[1]) {
			outputFile = reOutput.FindStringSubmatch(args[0])[1]
			alignement = reAlign.FindStringSubmatch(args[1])[1]
			userText = args[2]
		} else if reAlign.MatchString(args[0]) && reOutput.MatchString(args[1]) {
			alignement = reAlign.FindStringSubmatch(args[0])[1]
			outputFile = reOutput.FindStringSubmatch(args[1])[1]
			userText = args[2]		
		} else if reOutput.MatchString(args[0]) && !strings.HasPrefix(args[1], "--align") {
			outputFile = reOutput.FindStringSubmatch(args[0])[1]
			userText = args[1]	
			font = args[2]
		} else if reAlign.MatchString(args[0]) && !strings.HasPrefix(args[1], "--output") {
			alignement = reAlign.FindStringSubmatch(args[0])[1]
			userText = args[1]	
			font = args[2]
		} else {
			fmt.Println(alignErr)
			quit = true
		}
	case 4:
		if reOutput.MatchString(args[0]) && reAlign.MatchString(args[1]){
			outputFile = reOutput.FindStringSubmatch(args[0])[1]
			alignement = reAlign.FindStringSubmatch(args[1])[1]
			userText = args[2]	
			font = args[3]
		} else if reOutput.MatchString(args[1]) && reAlign.MatchString(args[0]){
			alignement = reAlign.FindStringSubmatch(args[0])[1]
			outputFile = reOutput.FindStringSubmatch(args[1])[1]
			userText = args[2]	
			font = args[3]			
		} else {
			fmt.Println(alignErr)
			quit = true
		}
	default:
		fmt.Println(alignErr)
		quit = true
	}
	return userText, font, outputFile, alignement, quit
}
