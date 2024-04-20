package main

import (
	"os"
	"formatTex/format"
	"fmt"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Error usage: <input.txt> <output.txt>")
		return
	}

	textBin, err := os.ReadFile(args[0]) 
	if err != nil {
		os.Stdout.WriteString("Error reading file: " + args[0])
		return
	}

	text := string(textBin) // Here we have our text as a string

	text = format.FixWhenFlagLast(text)
	text = format.Punctuation(text)
	//text = format.Format1(text)
	//text = format.Format2(text)
	

	fmt.Println(text)
}
