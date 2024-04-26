package main

import (
	"fmt"
	"formatTex/format"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		os.Stdout.WriteString("Error usage: <input.txt> <output.txt>\n")
		return
	}

	textBin, err := os.ReadFile(args[0])
	if err != nil {
		os.Stdout.WriteString("Error reading file: " + args[0])
		return
	}

	text := string(textBin) // Here we have our text as a string

	text = format.FlagsUserReact(text)
	text = format.Punctuation(text)
	text = format.Flags(text)
	text = format.Apostrophe(text)
	text = format.BasicGrammar(text)	
	text = format.Punctuation(text)
	text = format.CleanText(text)

	// Create a new file or truncate the existing file
	file, err := os.Create(args[1])
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
