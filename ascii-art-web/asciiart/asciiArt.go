package asciiart

import (
	"embed"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//go:embed banners
var Banners embed.FS

var fontLines int
var BadUserFont bool

// ASCIIArt function generate Ascii Art.
func ASCIIArt(userText, banner string) (string, error) {
	if GetAsciiTemplateByte(banner) == nil {
		if BadUserFont {
			return "only imported 8-lines fonts are supported", fmt.Errorf("invalid banner")
		}
		return "", fmt.Errorf("invalid banner")
	}
	if len(userText) == 0 {
		return "", nil
	}
	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if (asciiIndex < 32 || asciiIndex >= 127) && asciiIndex != 10 && asciiIndex != 13 {
			return "Non-ASCII characters aren't supported.\n", nil
		}
	}
	re := regexp.MustCompile(`\A((\\n)*)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")

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
			if char == '\n' || char == '\r' {
				continue
			}
			output += asciiTable[int(char-32)][i]

		}
		output += "\n"
	}
	return output
}

// This function check if a font is available at ./banners and at ../banners and lastely at ../fonts if found return its content as a slice of bytes else return nil.
func GetAsciiTemplateByte(font string) []byte {
	InitFontLines(font)
	asciiTemplateByte, err := Banners.ReadFile("banners/" + font + ".txt")
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
	default:
		fontLines = 8
	}
}
