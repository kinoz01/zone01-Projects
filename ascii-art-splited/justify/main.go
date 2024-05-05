package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

var tamara bool

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
		return
	}
	// these lines (20-->26) handle the cases of just new lines in the text.
	if userText == `\n` {
		fmt.Print("\n")
		return
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")

	var err error
	var asciiTemplateByte []byte

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
	if err != nil {
		fmt.Println("Error getting terminal width.")
		return
	}

	// Split asciiTemplate by double newline ("\n\n") to get individual ASCII characters from standard.txt.
	asciiCharacters := strings.Split(string(asciiTemplateByte), "\n\n")

	// Initialize asciiTable (2D table) (using "make" to avoid out of range).
	asciiTable := make([][]string, len(asciiCharacters))

	// Populate asciiTable [["1 line of A"...."8th line of A"]["1 line of B"...."8th line of B"]["1 line of C"...."8th line of C"]...["1 line of ~"..."8th line of ~"]].
	for i := range asciiCharacters {
		lines := strings.Split(asciiCharacters[i], "\n")
		asciiTable[i] = append(asciiTable[i], lines...)
	}

	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= len(asciiTable) {
			fmt.Println("🚨 Found an Invalid Ascii Character.")
			return
		}
	}

	switch alignement {
	case "left":
		fmt.Print(PrintAsciiArt(userText, "left", asciiTable, terminalWidth))
		return
	case "right":
		fmt.Print(PrintAsciiArt(userText, "right", asciiTable, terminalWidth))
		return
	case "center":
		fmt.Print(PrintAsciiArt(userText, "center", asciiTable, terminalWidth))
		return
	case "justify": 
		tamara = true
	}

	fmt.Print(PrintAsciiArt(userText, "left", asciiTable, terminalWidth))
	
	// fmt.Println(strings.Split(userText, `\n`))  // Printing the splited user text for clarification.
}

func PrintAsciiArt(userText, alignement string, asciiTable [][]string, terminalWidth int) string {
	var AsciiArt string
	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" {
			AsciiArt += "\n"
			continue
		}
		lenAscii := AsciiLineLEN(userLine, asciiTable)
		AsciiArt += PrintAsciiLine(userLine, alignement, asciiTable, lenAscii, terminalWidth)
	}
	return AsciiArt
}

func PrintAsciiLine(userLine, alignement string, asciiTable [][]string, lenAscii, terminalWidth int) string {
	var output string
	var row string
	for i := 0; i < 8; i++ {
		switch alignement{
		case "left":
			row = ""
		case "center":
			row = CenterSpaces(terminalWidth, lenAscii)
		case "right":
			row = RightSpaces(terminalWidth, lenAscii)
		}
		for _, char := range userLine {
			if char == ' ' && tamara {
				row += getJustifySpace(terminalWidth, userLine, asciiTable)
				continue
			}
			row += asciiTable[int(char-32)][i]
		}
		row += "\n"
		output += row
	}
	return output
}

func AsciiLineLEN(userLine string, asciiTable [][]string) int {
	if userLine == "" {
		return 0
	}
	var output string
	for i := 0; i < 8; i++ {
		row := ""
		for _, char := range userLine {
			row += asciiTable[int(char-32)][i]
		}
		row += "\n"
		output += row
	}
	outputSlice := strings.Split(output, "\n")
	return len(outputSlice[0])
}

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

func CenterSpaces(terminalWidth, lenAscii int) string {
	var spaces string
	spacesNum := (terminalWidth - lenAscii)/2
	for i := 0; i< spacesNum; i++ {
		spaces += " "
	}
	return spaces
}

func RightSpaces(terminalWidth, outputLen int) string {
	var spaces string
	spacesNum := (terminalWidth - outputLen)
	for i := 0; i< spacesNum; i++ {
		spaces += " "
	}
	return spaces
}

func getJustifySpace(terminalWidth int, userLine string, asciiTable [][]string) string {
	userWords := strings.Split(userLine, " ")
	var LenOfWords int
	var JustifySpace string
	for _, userWord := range userWords {
		LenOfWords += AsciiLineLEN(userWord, asciiTable)
	}
	JustifySpaceWidth := (terminalWidth - LenOfWords)/(len(userWords)-1)

	for j := 0; j < JustifySpaceWidth; j++ {
        JustifySpace += " "
    }
	return JustifySpace
}
