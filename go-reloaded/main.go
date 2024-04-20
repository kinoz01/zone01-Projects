package main

import (
	"fmt"
	"os"
	"formatTex/format"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Error usage: <input.txt> <output.txt>")
		return
	}

	textBin, err := os.ReadFile(args[0]) 
	if err != nil {
		fmt.Printf("Error reading file: %s\n", args[0])
		return
	}

	text := string(textBin) // Here we have our text as a string

	//text1 := format.Format1(text)
	//text1 := format.Format1(text)
	//text2 := format.Format1(text)
	text3 := format.Punctuation(text)
	fmt.Println(text3)
}
