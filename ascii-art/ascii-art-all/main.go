package main

import (
	"fmt"
	"os"
	"time"

	"asciiArt/asciiart"
)

func main() {

	args := os.Args[1:]
	userText, font, alignment, reverseInput, outputFiles, quit := asciiart.UserArgs(args)
	if quit {
		return
	}
	if reverseInput != "" {
		fmt.Print(asciiart.ReverseArt(reverseInput))
		return
	}
	userText, quit = asciiart.GetPrePrint(userText, outputFiles)
	if quit {
		fmt.Print(userText)
		return
	}
	terminalWidth, err := asciiart.GetTerminalWidth()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	output := asciiart.PrintAsciiArt(userText, alignment, asciiart.GetAsciiTable(font), terminalWidth)

	if outputFiles == nil {
		switch font {
		case "zigzag", "o2", "impossible", "univers":
			for _, char := range output {
				fmt.Print(string(char))
				time.Sleep(3 * time.Second / time.Duration(len(output))) // just some printing "art" for this particular fonts.
			}
		default:
			fmt.Print(output)
		}
	} else {
		for _, outputFile := range outputFiles {
			asciiart.GetAsciiFile(output, outputFile)
		}		
	}
}
