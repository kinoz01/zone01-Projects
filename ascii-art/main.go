package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

func main() {
/***********************************************************User Arguments***************************************************************/
	// This lines handle user arguments.
	args := os.Args[1:]
	var userText string
	var alignement string
	var font string
	reAlign := regexp.MustCompile(`\A--align=(center|left|right|justify)$`)
	// reColor := regexp.MustCompile(`\A--color=(\S+) +(\S+)$`)

	switch len(args) {
	case 0:
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	case 1:
		userText = args[0]
		font = "standard"
	case 2:
		if reAlign.MatchString(args[0]) {
			userText = args[1]
			font = "standard"
			alignement = reAlign.FindStringSubmatch(args[0])[1]
		} else {
			userText = args[0]
			font = args[1]
			if font != "standard" && font != "shadow" && font != "thinkertoy" {
				fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
				return
			}
		}
	case 3:
		if reAlign.MatchString(args[0]) {
			userText = args[1]
			alignement = reAlign.FindStringSubmatch(args[0])[1]
			font = args[2]
			if font != "standard" && font != "shadow" && font != "thinkertoy" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				return
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			return
		}
	}
/**************************************************************************************************************************************/

	if len(userText) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
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
	var outputLen int

	switch font {
	case "standard":
		asciiTemplateByte, err = os.ReadFile("./banners/standard.txt")
	case "shadow":
		asciiTemplateByte, err = os.ReadFile("./banners/shadow.txt")
	case "thinkertoy":
		asciiTemplateByte, err = os.ReadFile("./banners/thinkertoy.txt")
	}
	if err != nil {
		os.Stdout.WriteString("Error reading file: " + err.Error() + "\n")
		return
	}
	
	terminalWidth, err := getTerminalWidth()
	if err!= nil{
		fmt.Println("Error getting terminal width.")
		return
	}

	asciiTemplate := strings.ReplaceAll(string(asciiTemplateByte), "\r", "")

	// Split asciiTemplate by double newline ("\n\n") to get individual ASCII characters from standard.txt.
	asciiCharacters := strings.Split(asciiTemplate, "\n\n")

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
		if asciiIndex-32 < 0 || asciiIndex-32 > len(asciiTable) {
			fmt.Println("🚨 Found an Invalid Ascii Character.")
			return
		}
	}

	asciiOutput := PrintAscii(userText, asciiTable)
	sliceAsciioutput := strings.Split(asciiOutput, "\n")
	outputLen = len(sliceAsciioutput[0])
	// fmt.Println(outputLen)

	switch alignement {
	case "left":
		userText = AlignUserText(userText, "left", terminalWidth, outputLen)
	case "right":
		userText = AlignUserText(userText, "right", terminalWidth, outputLen)
	case "center":
		userText = AlignUserText(userText, "center", terminalWidth, outputLen)
	case "justify":
		userText = AlignUserText(userText, "justify", terminalWidth, outputLen)
	}



	asciiOutput = PrintAscii(userText, asciiTable)

	// Printing user input.
	fmt.Print(asciiOutput)
	// fmt.Println(strings.Split(userText, `\n`))  // Printing the splited user text for clarification.
}

func PrintAscii(userText string, asciiTable [][]string) string {
	var asciiOutput string
	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" { // result of spliting.
			asciiOutput += "\n"
			continue
		}
		asciiOutput += PrintAsciiWord(userLine, asciiTable)
	}
	return asciiOutput
}

func PrintAsciiWord(userLine string, asciiTable [][]string) string {
	var asciiWord string
	for i := 0; i < 8; i++ {
		for _, userTextChar := range userLine {
			asciiIndex := int(userTextChar)
			asciiWord += asciiTable[asciiIndex-32][i]
		}
		asciiWord += "\n"
	}
	return asciiWord
}

/********** How did I come up with the printing mechanism? *************/
// asciiTable[32][0] + " " +  asciiTable[33][0] + "\n" + asciiTable[32][1] + " " +  asciiTable[33][1] + "\n" + asciiTable[32][2] + " " +  asciiTable[33][2] + "\n" ....ect
// Just by trying these!

// Function to get the current terminal width.
func getTerminalWidth() (int, error) {
	var dimensions struct {
		Rows uint16
		Cols uint16
	}

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdout), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)))
	if err != 0 {
		return 0, err
	}

	return int(dimensions.Cols), nil
}

func AlignUserText(userText, alignement string, terminalWidth, outputLen int) string {
	if alignement == "left" {
		return userText
	} else if alignement == "center" {
		return CenterSpaces(terminalWidth, outputLen) + userText + CenterSpaces(terminalWidth, outputLen)
	} else if alignement == "right" {
		return RightSpaces(terminalWidth, outputLen) + userText
	}
	return userText
}

func CenterSpaces(terminalWidth, outputLen int) string {
	var spaces string
	spacesNum := (terminalWidth - outputLen)/2
	for i := 0; i< spacesNum/6; i++ {
		spaces += " "
	}
	return spaces
}

func RightSpaces(terminalWidth, outputLen int) string {
	var spaces string
	spacesNum := (terminalWidth - outputLen)
	for i := 0; i< spacesNum/6; i++ {
		spaces += " "
	}
	return spaces
}