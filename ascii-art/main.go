package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	/***************************************/
	// This lines handle user arguments.
	args := os.Args[1:]
	if len(args) <= 0 || len(args) > 2 || (len(args) == 2 && args[1] != "standard" && args[1] != "shadow" && args[1] != "thinkertoy") {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	userText := args[0]
	if len(userText) == 0 {
		return
	}

	/*-----------------------------------------------------------------------------------*/
	// These Lines (20-->26) handle the cases of just new lines ("\n\n...") in the text.
	/*-----------------------------------------------------------------------------------*/
	if userText == `\n` {
		fmt.Print("\n")
		return
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")

	var err error
	var asciiTemplateByte []byte
	if len(args) == 2 {
		switch args[1] {
		case "standard":
			asciiTemplateByte, err = os.ReadFile("./banners/standard.txt")
		case "shadow":
			asciiTemplateByte, err = os.ReadFile("./banners/shadow.txt")
		case "thinkertoy":
			asciiTemplateByte, err = os.ReadFile("./banners/thinkertoy.txt")
		}
	} else {
		asciiTemplateByte, err = os.ReadFile("./banners/standard.txt")
	}
	if err != nil {
		os.Stdout.WriteString("Error reading file: " + err.Error() + "\n")
		return
	}

	asciiTemplate := strings.ReplaceAll(string(asciiTemplateByte), "\r", "")

	// Split asciiTemplate by double newline ("\n\n") to get individual ASCII characters from standard.txt.
	asciiCharacters  := strings.Split(asciiTemplate, "\n\n")

	// Initialize asciiTable (2D table) (using "make" to avoid out of range).
	asciiTable := make([][]string, len(asciiCharacters))

	// Populate asciiTable [["1 line of A"...."8th line of A"]["1 line of B"...."8th line of B"]["1 line of C"...."8th line of C"]...["1 line of ~"..."8th line of ~"]].
	for i := range asciiCharacters {
		lines := strings.Split(asciiCharacters[i], "\n")
		asciiTable[i] = append(asciiTable[i], lines...)
	}

	/*------------------------------------------------------------------------------------------------------------------------------------*/
	// This loop check user input searching for invalid ascii and returning if found any to avoid out of range when invalid ascii in input.
	/*------------------------------------------------------------------------------------------------------------------------------------*/
	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= len(asciiTable) {
			fmt.Println("🚨 Found an Invalid Ascii Character.") 
			return
		}
	}

	var asciiOutput string

	// Printing user input.
	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" { // result of spliting.
			asciiOutput += "\n"
			continue
		}
		asciiOutput += PrintAscii(userLine, asciiTable)
	}

	fmt.Print(asciiOutput)
	// fmt.Println(strings.Split(userText, `\n`))  // Printing the splited user text for clarification.
}

func PrintAscii(userLine string, asciiTable [][]string) string {
	var asciiOutput string
	for i := 0; i < 8; i++ {
		for _, userTextChar := range userLine {
			asciiIndex := int(userTextChar)
			asciiOutput += asciiTable[asciiIndex-32][i]
		}
		asciiOutput += "\n"
	}
	return asciiOutput
}

// Function to get the current terminal width.
// func getTerminalWidth() (int, error) {
// 	var dimensions struct {
// 		Rows uint16
// 		Cols uint16
// 	}

// 	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdout), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)))
// 	if err != 0 {
// 		return 0, err
// 	}

// 	return int(dimensions.Cols), nil
// }

/********** How did I come up with the printing mechanism? *************/
// asciiTable[32][0] + " " +  asciiTable[33][0] + "\n" + asciiTable[32][1] + " " +  asciiTable[33][1] + "\n" + asciiTable[32][2] + " " +  asciiTable[33][2] + "\n" ....ect
// Just by trying these!
