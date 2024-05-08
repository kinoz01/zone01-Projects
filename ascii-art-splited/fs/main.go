package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) > 2 || len(args) == 0 {
		os.Stdout.WriteString("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\n")
		return
	}
	userText := args[0]
	if len(userText) == 0 {
		return
	}
	// these lines (20-->26) handle the cases of just new lines in the text.
	if userText == `\n` {
		fmt.Print("\n")
		return
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")

	var asciiTemplateByte []byte
	var err error
	if len(args) == 1 {
		asciiTemplateByte, err = os.ReadFile("./banners/standard.txt")
	} else {
		asciiTemplateByte, err = os.ReadFile("./banners/" + args[1] + ".txt")
	}
	if err != nil {
		os.Stdout.WriteString("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\n")
		return
	}
	asciiTemplate := strings.ReplaceAll(string(asciiTemplateByte), "\r", "")

	// Split asciiTemplate by double newline ("\n\n") to get individual ASCII characters from standard.txt.
	asciiCharacters := strings.Split(string(asciiTemplate[1:]), "\n\n")

	// Initialize asciiTable (2D table) (using "make" to avoid out of range).
	asciiTable := make([][]string, len(asciiCharacters))

	// Populate asciiTable [["1 line of A"...."8th line of A"]["1 line of B"...."8th line of B"]["1 line of C"...."8th line of C"]...["1 line of ~"..."8th line of ~"]].
	for i := range asciiCharacters {
		lines := strings.Split(asciiCharacters[i], "\n")
		asciiTable[i] = append(asciiTable[i], lines...)
	}

	// Searching for invalid ascii to avoid out of range panic.
	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= len(asciiTable) {
			fmt.Println("🚨 Found an Invalid Ascii Character.")
			return
		}
	}

	// Printing mechanism.
	var output string
	userLine := strings.Split(userText, `\n`)
	for _, newLine := range userLine {
		if newLine == "" {
			output += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range newLine {
				output += asciiTable[int(char)-32][i]
			}
			output += "\n"
		}
	}
	if len(args) == 2 && args[1] == "zigzag" {
		for _, char := range output {
			fmt.Print(string(char))
			time.Sleep(3 * time.Second / time.Duration(len(output))) // just some printing "art" for this particular font.
		}
	} else {
		fmt.Print(output)
	}	
}
