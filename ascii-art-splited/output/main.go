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
		fmt.Println(userText)
		return
	}
	output := asciiart.Print(userText, font)
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
