package main

import (
	"asciiArt/asciiart"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	userText, font, outputFile := asciiart.UerArgs(args)
	if userText == "" {
		return
	}
	userText, quit := asciiart.PrePrint(userText)
	if quit {
		fmt.Print(userText)
		return
	}
	output := asciiart.Print(userText, font)
	
	if outputFile == "" {
		switch font {
		case "zigzag":
			for _, char := range output {
				fmt.Print(string(char))
				time.Sleep(3 * time.Second / time.Duration(len(output))) // just some printing "art" for this particular fonts.							
			}
		default:
			fmt.Print(output)
		}
	} else {
		asciiart.CreateFile(output, outputFile)
	}
}
