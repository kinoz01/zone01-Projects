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

	tests := []struct {
		name string // Name of the test case
		text string // Input text
		want string // Expected output using a regexp pattern
	}{
		{"Test 1", `\n`, "\n"},
		{"Test 2", `\n\n`, "\n\n"},
		{"Test 3", "hello", string(want1)},
		{"Test 4", "HELLO", string(want2)},
		{"Test 5", "HeLlo HuMaN", string(want3)},
		{"Test 6", "1Hello 2There", string(want4)},
		{"Test 7", `Hello\nThere`, string(want5)},
		{"Test 8", `Hello\n\nThere`, string(want6)},
		{"Test 9", `{Hello & There #}`, string(want7)},
		{"Test 10", `hello There 1 to 2!`, string(want8)},
		{"Test 11", `MaD3IrA&LiSboN`, string(want9)},
		{"Test 12", "1a\"#FdwHywR&/()=", string(want10)},
		{"Test 13", "{|}~", string(want11)},
		{"Test 14", `[\]^_ 'a`, string(want12)},
		{"Test 15", "RGB", string(want13)},
		{"Test 16", ":;<=>?@", string(want14)},
		{"Test 17", `\!" #$%&` + "'" + `()*+,-./`, string(want15)},
		{"Test 18", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(want16)},
		{"Test 19", "abcdefghijklmnopqrstuvwxyz", string(want17)},
		{"Test Error", "²", "🚨 Found an Invalid Ascii Character."},
	}
	
	for _, cas := range tests {
		t.Run(cas.name, func(t *testing.T) {
			// Compile the expected regexp.
			// Call the function.
			got := AsciiArt(cas.text)
			// Check if the output matches the expected regexp.
			if got != cas.want {
				t.Errorf("For input %q, expected match for %#q but got %q", cas.text, cas.want, got)
			}
		})
	}
}

func AsciiArt(userText string) string {
	if len(userText) == 0 {
		os.Exit(1)
	}
	if userText == `\n` {
		return "\n"
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")

	asciiTemplateByte, err := os.ReadFile("../banners/standard.txt")
	if err != nil {
		os.Stdout.WriteString("Error reading file: standard.txt\n")
		os.Exit(1)
	}

	// Split asciiTemplate by double newline ("\n\n") to get individual ASCII characters.
	asciiCharacters := strings.Split(string(asciiTemplateByte), "\n\n")

	// Initialize asciiTable (2D table) (using "make" to avoid out of range).
	asciiTable := make([][]string, len(asciiCharacters))

	for _ , userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= len(asciiTable) {
			return"🚨 Found an Invalid Ascii Character."
		}
	}

	// Populate asciiTable [["1 line of A"...."8th line of A"]["1 line of B"...."8th line of B"]["1 line of C"...."8th line of C"]...["1 line of ~"..."8th line of ~"]].
	for i := range asciiCharacters {
		lines := strings.Split(asciiCharacters[i], "\n")
		asciiTable[i] = append(asciiTable[i], lines...)
	}

	var AsciiArt string
	// printing user input.
	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" {
			AsciiArt += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _ , userTextChar := range userLine {
				asciiIndex := int(userTextChar)
				AsciiArt += asciiTable[asciiIndex-32][i]
			}
			AsciiArt +="\n"
		}
	}
	return AsciiArt
}
