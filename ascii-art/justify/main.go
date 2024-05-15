package main

import (
	"asciiArt/asciiart"
	"fmt"
	"os"
	"time"
)


func main() {
	args := os.Args[1:]
	userText, font, outputFile, alignement, quit := asciiart.UserArgs(args)
	fmt.Println(quit)
	if userText == "" || quit{
		return
	}
	userText, quit = asciiart.PrePrint(userText)
	if quit {
		fmt.Print(userText)
		return
	}
	terminalWidth, err := asciiart.GetTerminalWidth()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	output := asciiart.PrintAsciiArt(userText, alignement, asciiart.GetAsciiTable(font), terminalWidth)

	if outputFile == "" {
		switch font {
		case "zigzag", "o2":
			for _, char := range output {
				fmt.Print(string(char))
				time.Sleep(3 * time.Second / time.Duration(len(output))) // just some printing "art" for this particular font.							
			}
		default:
			fmt.Print(output)
		}
	} else {
		asciiart.CreateFile(output, outputFile)
	}
}
