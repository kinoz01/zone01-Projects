package asciiart

import (
	"embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed banners
var banners embed.FS

var BadUserFont bool

var fontLines int

// ASCIIArt function generate Ascii Art.
func ASCIIArt(userText, banner string) (string, error) {
	msg, quit := GetPrePrint(userText, banner)
	if quit {
		return msg, fmt.Errorf(msg)
	}
	var AsciiArt string
	for _, userLine := range strings.Split(userText, "\r\n") {
		if userLine == "" {
			AsciiArt += "\n"
			continue
		}
		AsciiArt += PrintAsciiLine(userLine, GetAsciiTable(banner))
	}
	return AsciiArt, nil
}

func PrintAsciiLine(userLine string, asciiTable [][]string) string {
	var output string
	for i := 0; i < fontLines; i++ {
		for _, char := range userLine {
			output += asciiTable[int(char-32)][i]

		}
		output += "\n"
	}
	return output
}

// This function check if a font is available at ./banners and at ../banners and lastely at ../fonts if found return its content as a slice of bytes else return nil.
func GetAsciiTemplateByte(font string) []byte {
	InitFontLines(font)
	asciiTemplateByte, err := banners.ReadFile("banners/" + font + ".txt")
	if err != nil {
		return nil
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

// this function initialize the number of lines present in the ascii character so we can range on them later.
func InitFontLines(font string) {
	switch font {
	case "graceful":
		fontLines = 4
	case "small":
		fontLines = 5
	case "phoenix", "o2", "starwar", "stop", "varsity":
		fontLines = 7
	case "standard", "shadow", "thinkertoy", "arob", "zigzag", "henry3D", "doom", "tiles", "jacky", "catwalk", "coins":
		fontLines = 8
	case "fire":
		fontLines = 9
	case "jazmine", "matrix":
		fontLines = 10
	case "blocks", "univers":
		fontLines = 11
	case "impossible":
		fontLines = 12
	case "georgi":
		fontLines = 16
	}
}

// Quit in case of empty string "", search for invalid ascii to avoid out of range panic, and remove one new line in case of just new lines in userInput.
func GetPrePrint(userText, banner string) (string, bool) {
	if GetAsciiTemplateByte(banner) == nil {
		return "🚨 Invalid Banner.\n", true
	}
	if len(userText) == 0 {
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
