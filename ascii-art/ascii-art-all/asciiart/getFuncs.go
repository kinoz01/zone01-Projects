package asciiart

import (
	"embed"
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

//go:embed banners
var banners embed.FS

var BadUserFont bool

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

// This function check if a font is available at ./banners (both in built and go run modes) if found return its content as a slice of bytes else return nil.
func GetAsciiTemplateByte(font string) []byte {
	InitFontLines(font)
	asciiTemplateByte, err := banners.ReadFile("banners/" + font + ".txt")
	// fs.ReadFile to read embedded floder both in build and run mode.
	// asciiTemplateByte, err := fs.ReadFile(banners, "banners/"+font+".txt")
	if err != nil {
		// in case of builded program we check banners folder.
		asciiTemplateByte, err = os.ReadFile("./banners/" + font + ".txt")
		if err != nil {
			return nil
		}
		// if we find font but it's unsupportable.
		// this will never be reached in case of "go run ." because the place we read from don't change---->"./banners/" contrary when built.
		if len(strings.Split(string(asciiTemplateByte), "\n")) != 856 || regexp.MustCompile(`\n\n\n`).MatchString(string(asciiTemplateByte)) {
			BadUserFont = true
			return nil
		}
	}
	return asciiTemplateByte
}

// Quit in case of empty string "", search for invalid ascii to avoid out of range panic, and remove one new line in case of just new lines in userInput.
func GetPrePrint(userText string, outputFiles []string) (string, bool) {
	if len(userText) == 0 {
		for _, outputFile := range outputFiles {
			GetAsciiFile("", outputFile)
		}
		return "", true
	}
	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= 95 {
			return "🚨 Found an Invalid Ascii Character.\n", true
		}
	}
	re := regexp.MustCompile(`\A((\\n)*)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")
	return userText, false
}

// Get the current terminal width using syscall request to the kernel.
func GetTerminalWidth() (int, error) {
	var dimensions [4]uint16
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(2), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&dimensions)))
	if err != 0 {
		return 0, err
	}
	return int(dimensions[1]), nil
}

// Get the len of the first line of the printed ascii before justification/alignement.
func GetAsciiLineLen(userLine string, asciiTable [][]string) int {
	var output string
	for _, char := range userLine {
		output += asciiTable[int(char-32)][0]
	}
	return len([]rune(output)) // converting to []rune in case font contains special characters like "zigzag".
}

// Get Spaces to be placed AT EACH user input text SPACE for justification.
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

// Modify the slice of colors, that will be used while printing
func GetColorSlice(color, colorChars, userText string) {
	indices := []int{}
	if !strings.Contains(userText, colorChars) {
		return
	} else {
		for i := 0; i < len(userText)-len(colorChars)+1; i++ {
			if userText[i:i+len(colorChars)] == colorChars {
				for j := i; j < i+len(colorChars); j++ {
					indices = append(indices, j) // just the normal index function but we add this loop to get the indice of each character.
				}
				i += len(colorChars) - 1 // added to fix --color=red ll llllllop
			}
		}
	}
	for _, indexToColor := range indices {
		ColorSlice[indexToColor] = color
	}
}

// This func is to solve the go test errors when there is no colors.
func ColorSliceEmpty() bool {
	for _, str := range ColorSlice {
		if str != "" {
			return false
		}
	}
	return true
}

// takes the content and the file name, create the file and write the content (ascii art) in it.
func GetAsciiFile(output, outputFileName string) {
	err := os.WriteFile(outputFileName, []byte(output), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
