package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"  
)

func main() {
	var outputAsciiFile string
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Enter a valid argument")
	}
	asciiTemplateByte, err := os.ReadFile("./" + args[0])
	if err != nil {
		fmt.Println("Error reading file.")
		return
	}
	
	outputAsciiFile = strings.ReplaceAll(string(asciiTemplateByte), "@@", "\n")
	//outputAsciiFile = strings.ReplaceAll(string(asciiTemplateByte), "##", "\n")
	outputAsciiFile = strings.ReplaceAll(outputAsciiFile, "$", " ")
	//outputAsciiFile = strings.ReplaceAll(outputAsciiFile, "_", " ")

	re := regexp.MustCompile(`(#|@)`)
	outputAsciiFile = re.ReplaceAllString(outputAsciiFile, "")


	CreateFile(outputAsciiFile, args[1]+".txt")
}

func CreateFile(output, outputFile string) {
	// Create a new file or truncate the existing file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
