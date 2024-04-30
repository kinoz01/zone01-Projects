// v1.6
package main

import (
	"fmt"
	"formatTex/format"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		os.Stdout.WriteString("❌ Please enter a valid number of arguments. Usage: \"input.txt output.txt\"\n")
		return
	}

	textBin, err := os.ReadFile(args[0])
	if err != nil {
		os.Stdout.WriteString("Error reading file: " + args[0] + "\n")
		return
	}

	if strings.HasSuffix(args[1], ".go")  {
		panic("HEY what are you doing bro?")
	}
	
	text := string(textBin) // Here we have our text as a string

	text = format.FlagsWrongUsage(text)
	text = format.Flags(text)
	text = format.Apostrophe(text)
	text = format.BasicGrammar(text)	
	text = format.Punctuation(text)
	text = format.CleanText(text)
	text += "\n"  // for the cat command  

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
