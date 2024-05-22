package asciiart

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// This function check if a font is available at ./banners and at ../banners and lastely at ../fonts if found return its content as a slice of bytes else return nil.
func GetAsciiTemplateByte(font string) []byte {
	InitFontLines(font)

	asciiTemplateByte, err := os.ReadFile("./banners/" + font + ".txt")
	if err != nil {
		// If reading with fails try checking if there are any outside/user fonts in a folder called banners along side the excutable program.
		asciiTemplateByte, err = os.ReadFile("../banners/" + font + ".txt")
		if err != nil {
			asciiTemplateByte, err = os.ReadFile("../fonts/" + font + ".txt")
			if err != nil {
				// If all three attempts fail return.
				return nil
			}
		}
	}
	return asciiTemplateByte
}

// Get an ascii table by transforming the font text from (file name------> slice of bytes------> slice of string (splited by \n\n)--------> 2D table (ready to print)).
func GetAsciiTable(font string) [][]string {
	asciiTemplate := strings.ReplaceAll(string(GetAsciiTemplateByte(font)), "\r", "")
	asciiCharacters := strings.Split(string(asciiTemplate[1:]), "\n\n")
	asciiTable := make([][]string, len(asciiCharacters))

	for i := range asciiCharacters {
		lines := strings.Split(asciiCharacters[i], "\n")
		asciiTable[i] = append(asciiTable[i], lines...)
	}
	return asciiTable
}

// Get the current terminal width.
func GetTerminalWidth() (int, error) {
	var dimensions [2]uint16
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(2), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&dimensions)))
	if err != 0 {
		return 0, err
	}
	return int(dimensions[1]), nil
}

// Get the len of the first line of the printed ascii before justification/alignement.
func GetAsciiLineLen(userLine string, asciiTable [][]string) int {
	if userLine == "" {
		return 0
	}
	var output string
	for _, char := range userLine {
		output += asciiTable[int(char-32)][0]
	}
	return len([]rune(output)) // converting to []rune in case font contains special characters like "zigzag".
}

//Get Spaces to be placed AT EACH user input text SPACE for justification.
func GetJustifySpace(terminalWidth int, userLine string, asciiTable [][]string) string {
	userWords := strings.Split(userLine, " ")
	var LenOfWords int
	for _, userWord := range userWords {
		LenOfWords += GetAsciiLineLen(userWord, asciiTable)
	}
	if terminalWidth-LenOfWords > 0 { // remove out of range when len of the printed text is bigger than terminal (Ex: go run . --color=ocean 01 --align=right "zone 01 Oujda" impossible).
		return strings.Repeat(" ", (terminalWidth-LenOfWords)/(len(userWords)-1))
	}
	return ""
}

func JustifyOneWordSpaces(userLine string) string {
	if userLine[len(userLine)-1] == ' ' && userLine[0] == ' ' {
		userLine = " " + strings.Fields(userLine)[0] + " "
	} else if userLine[0] == ' ' {
		userLine = " " + strings.Fields(userLine)[0]
	} else if userLine[len(userLine)-1] == ' ' {
		userLine = strings.Fields(userLine)[0] + " "
	}
	return userLine
}
