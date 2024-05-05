package test

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	// Define test cases
	want1, _ := os.ReadFile("./cases/want1.txt")
	want2, _ := os.ReadFile("./cases/want2.txt")
	want3, _ := os.ReadFile("./cases/want3.txt")
	want4, _ := os.ReadFile("./cases/want4.txt")
	want5, _ := os.ReadFile("./cases/want5.txt")
	want6, _ := os.ReadFile("./cases/want6.txt")
	want7, _ := os.ReadFile("./cases/want7.txt")
	want8, _ := os.ReadFile("./cases/want8.txt")
	want9, _ := os.ReadFile("./cases/want9.txt")
	want10, _ := os.ReadFile("./cases/want10.txt")
	want11, _ := os.ReadFile("./cases/want11.txt")
	want12, _ := os.ReadFile("./cases/want12.txt")
	want13, _ := os.ReadFile("./cases/want13.txt")
	want14, _ := os.ReadFile("./cases/want14.txt")
	want15, _ := os.ReadFile("./cases/want15.txt")
	want16, _ := os.ReadFile("./cases/want16.txt")
	want17, _ := os.ReadFile("./cases/want17.txt")
	want18, _ := os.ReadFile("./cases/want18.txt")
	want19, _ := os.ReadFile("./cases/want19.txt")
	want20, _ := os.ReadFile("./cases/want20.txt")
	want21, _ := os.ReadFile("./cases/want21.txt")
	want22, _ := os.ReadFile("./cases/want22.txt")
	want23, _ := os.ReadFile("./cases/want23.txt")
	want24, _ := os.ReadFile("./cases/want24.txt")
	want25, _ := os.ReadFile("./cases/want25.txt")

	tests := []struct {
		name string // Name of the test case
		text string // Input text
		want string // Expected output using a regexp pattern
		banner string // choosed font
	}{
		{"Test 1", `\n`, "\n", "standard"},
		{"Test 2", `\n\n`, "\n\n", "standard"},
		{"Test 3", "hello", string(want1), "standard"},
		{"Test 4", "HELLO", string(want2), "standard"},
		{"Test 5", "HeLlo HuMaN", string(want3), "standard"},
		{"Test 6", "1Hello 2There", string(want4), "standard"},
		{"Test 7", `Hello\nThere`, string(want5), "standard"},
		{"Test 8", `Hello\n\nThere`, string(want6), "standard"},
		{"Test 9", `{Hello & There #}`, string(want7), "standard"},
		{"Test 10", `hello There 1 to 2!`, string(want8), "standard"},
		{"Test 11", `MaD3IrA&LiSboN`, string(want9), "standard"},
		{"Test 12", "1a\"#FdwHywR&/()=", string(want10), "standard"},
		{"Test 13", "{|}~", string(want11), "standard"},
		{"Test 14", `[\]^_ 'a`, string(want12), "standard"},
		{"Test 15", "RGB", string(want13), "standard"},
		{"Test 16", ":;<=>?@", string(want14), "standard"},
		{"Test 17", `\!" #$%&` + "'" + `()*+,-./`, string(want15), "standard"},
		{"Test 18", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(want16), "standard"},
		{"Test 19", "abcdefghijklmnopqrstuvwxyz", string(want17), "standard"},
		{"Test Error", "²", "Found an Invalid Ascii Character.", "standard"},

		{"Test 20", "hello world", string(want18), "shadow"},
		{"Test 21", "nice 2 meet you", string(want19), "thinkertoy"},
		{"Test 22", "you & me", string(want20), "standard"},
		{"Test 23", "123", string(want21), "shadow"},
		{"Test 24", "/(\")", string(want22), "thinkertoy"},
		{"Test 25", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(want23), "shadow"},
		{"Test 26", "\"#$%&/()*+,-./", string(want24), "thinkertoy"},
		{"Test 27", "It's Working", string(want25), "thinkertoy"},
	}
	
	for _, cas := range tests {
		t.Run(cas.name, func(t *testing.T) {
			// Compile the expected regexp.
			// Call the function.
			got := AsciiArt(cas.text, cas.banner)
			// Check if the output matches the expected regexp.
			if got != cas.want {
				t.Errorf("For input %q, expected match for %#q but got %q", cas.text, cas.want, got)
			}
		})
	}
}

// A function that takes user text and banner and return Ascii Art.
func AsciiArt(userText, banner string) string {
	if len(userText) == 0 {
		os.Exit(1)
	}
	if userText == `\n` {
		return "\n"
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")


	var err error
	var asciiTemplateByte []byte
	switch banner {
	case "standard":
		asciiTemplateByte, err = os.ReadFile("../banners/standard.txt")
	case "shadow":	
		asciiTemplateByte, err = os.ReadFile("../banners/shadow.txt")
	case "thinkertoy": 
		asciiTemplateByte, err = os.ReadFile("../banners/thinkertoy.txt")
	default:
		return 	"Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"
	}
	if err != nil {
		os.Stdout.WriteString("Error reading file: standard.txt\n")
		os.Exit(1)
	}
	asciiTemplate := strings.ReplaceAll(string(asciiTemplateByte), "\r", "")

	// Split asciiTemplate by double newline ("\n\n") to get individual ASCII characters.
	asciiCharacters := strings.Split(asciiTemplate, "\n\n")

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
			return "Found an Invalid Ascii Character."
		}
	}

	var AsciiArt string
	// printing user input.
	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" {
			AsciiArt += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, userTextChar := range userLine {
				asciiIndex := int(userTextChar)
				AsciiArt += asciiTable[asciiIndex-32][i]
			}
			AsciiArt += "\n"
		}
	}
	return AsciiArt
}
