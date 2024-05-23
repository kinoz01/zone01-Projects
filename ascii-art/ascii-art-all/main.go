package main

import (
	"asciiArt/asciiart"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	userText, font, alignment, outputFile, reverseInput, colorMap, quit := asciiart.UserArgs(args)
	if quit {
		return
	}

	if reverseInput != "" {
		fmt.Print(asciiart.ReverseArt(reverseInput))
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
	output := asciiart.PrintAsciiArt(userText, alignment, asciiart.GetAsciiTable(font), terminalWidth, colorMap)

	if outputFile == "" {
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
		asciiart.GetAsciiFile(output, outputFile)
	}
}
