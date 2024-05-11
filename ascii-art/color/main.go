package main

import (
	"asciiArt/asciiart"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	args, colorMap, quit := asciiart.GetColorMap(args)
	if quit {
		return
	}
	fmt.Println(colorMap)
	userText, font, outputFile, alignement := asciiart.UserArgs(args)
	if userText == "" {
		return
	}
	userText, quit = asciiart.PrePrint(userText)
	if quit {
		fmt.Print(userText)
		return
	}
	terminalWidth, err := asciiart.GetTerminalWidth()
	if err != nil {
		return
	}
	output := asciiart.PrintAsciiArt(userText, alignement, asciiart.GetAsciiTable(font), terminalWidth)

	if outputFile == "" {
		switch font {
		case "zigzag", "o2":
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
