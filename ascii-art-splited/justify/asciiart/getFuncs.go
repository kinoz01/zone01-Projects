package asciiart

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func GetAsciiTable(font string) [][]string {
	InitFontLines(font)
	asciiTemplateByte, err := os.ReadFile("./banners/" + font + ".txt")
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return nil
	}
	asciiTemplate := strings.ReplaceAll(string(asciiTemplateByte), "\r", "")
	asciiCharacters := strings.Split(string(asciiTemplate[1:]), "\n\n")
	asciiTable := make([][]string, len(asciiCharacters))

	for i := range asciiCharacters {
		lines := strings.Split(asciiCharacters[i], "\n")
		asciiTable[i] = append(asciiTable[i], lines...)
	}
	return asciiTable
}

// Function to get the current terminal width.
func GetTerminalWidth() (int, error) {
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

func GetAsciiLineLen(userLine string, asciiTable [][]string) int {
	if userLine == "" {
		return 0
	}
	var output string
	for i := 0; i < fontLines; i++ {
		for _, char := range userLine {
			output += asciiTable[int(char-32)][i]
		}
		output += "\n"
	}
	outputSlice := strings.Split(output, "\n")
	return len(outputSlice[0])
}

func GetCenterSpaces(terminalWidth, lenAscii int) string {
	var spaces string
	spacesNum := (terminalWidth - lenAscii)/2
	for i := 0; i < spacesNum; i++ {
		spaces += " "
	}
	return spaces
}

func GetRightSpaces(terminalWidth, outputLen int) string {
	var spaces string
	spacesNum := (terminalWidth - outputLen)
	for i := 0; i< spacesNum; i++ {
		spaces += " "
	}
	return spaces
}

func GetJustifySpace(terminalWidth int, userLine string, asciiTable [][]string) string {
	userWords := strings.Split(userLine, " ")
	var LenOfWords int
	var JustifySpace string
	for _, userWord := range userWords {
		LenOfWords += GetAsciiLineLen(userWord, asciiTable)
	}
	JustifySpaceWidth := (terminalWidth - LenOfWords)/(len(userWords)-1)

	for j := 0; j < JustifySpaceWidth; j++ {
        JustifySpace += " "
    }
	return JustifySpace
}
