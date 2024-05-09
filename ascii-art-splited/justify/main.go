package main

import (
	"asciiArt/asciiart"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	userText, font, outputFile, alignement := asciiart.UerArgs(args)
	if userText == "" {
		return
	}
	userText, quit := asciiart.PrePrint(userText)
	if quit {
		fmt.Print(userText)
		return
	}
	terminalWidth, err := asciiart.GetTerminalWidth()
	if err != nil {
		fmt.Println("Error getting terminal width.")
		return
	}
	
	output := asciiart.PrintAsciiArt(userText, alignement, asciiart.GetAsciiTable(font), terminalWidth)
	if font == "zigzag" {
		for _, char := range output {
			fmt.Print(string(char))
			time.Sleep(4 * time.Second / time.Duration(len(output))) // just some printing "art" for this particular font.
		}

	} else {
		fmt.Print(output)
	}
	if outputFile != "" {
		asciiart.CreateFile(output, outputFile)
	}
}
