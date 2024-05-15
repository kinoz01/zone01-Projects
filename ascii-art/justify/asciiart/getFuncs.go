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
		fmt.Println(alignErr)
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
	var dimensions [4]uint16
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stderr), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&dimensions)))
	if err != 0 {
		return 0, err
	}
	return int(dimensions[1]), nil
}

func GetAsciiLineLen(userLine string, asciiTable [][]string) int {
	var output string
	if userLine == "" {
		return 0
	}
	for _, char := range userLine {
		output += asciiTable[int(char-32)][0]
	}
	return len([]rune(output)) // converting to []rune in case font contains special characters like zigzag.
}

func GetJustifySpace(terminalWidth int, userLine string, asciiTable [][]string) string {
	userWords := strings.Split(userLine, " ")
	var LenOfWords int
	for _, userWord := range userWords {
		LenOfWords += GetAsciiLineLen(userWord, asciiTable)
	}
	return strings.Repeat(" ", (terminalWidth - LenOfWords) / (len(userWords) - 1))
}
