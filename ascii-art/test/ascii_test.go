package main

import (
	"os"
	"strings"
	"testing"
	"regexp"
)

func TestAsciiArt(t *testing.T) {
	// Define test cases
	var tests = []struct {
		name string // Name of the test case
		text string // Input text
		want string // Expected output using a regexp pattern
	}{
{"Test 1", 
`\n`, 
"\n"},
{"Test 2", 
"hello", 
` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`},
{"Test 3", 
"HELLO", 
` _    _   ______   _        _         ____   
| |  | | |  ____| | |      | |       / __ \  
| |__| | | |__    | |      | |      | |  | | 
|  __  | |  __|   | |      | |      | |  | | 
| |  | | | |____  | |____  | |____  | |__| | 
|_|  |_| |______| |______| |______|  \____/  
                                             
                                             
`},
{"Test 4", 
"HeLlo HuMaN", 
" _    _          _        _                 _    _           __  __           _   _  \n| |  | |        | |      | |               | |  | |         |  \\/  |         | \\ | | \n| |__| |   ___  | |      | |   ___         | |__| |  _   _  | \\  / |   __ _  |  \\| | \n|  __  |  / _ \\ | |      | |  / _ \\        |  __  | | | | | | |\\/| |  / _` | | . ` | \n| |  | | |  __/ | |____  | | | (_) |       | |  | | | |_| | | |  | | | (_| | | |\\  | \n|_|  |_|  \\___| |______| |_|  \\___/        |_|  |_|  \\__,_| |_|  |_|  \\__,_| |_| \\_| \n                                                                                     \n                                                                                     \n"},

{"Test 5", 
"1Hello 2There", 
"     _    _          _   _                         _______   _                           \n _  | |  | |        | | | |                ____   |__   __| | |                          \n/ | | |__| |   ___  | | | |   ___         |___ \\     | |    | |__     ___   _ __    ___  \n| | |  __  |  / _ \\ | | | |  / _ \\          __) |    | |    |  _ \\   / _ \\ | '__|  / _ \\ \n| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ \n|_| |_|  |_|  \\___| |_| |_|  \\___/        |_____|    |_|    |_| |_|  \\___| |_|     \\___| \n                                                                                         \n                                                                                         \n"},
{"Test 6", 
"HELLO", 
` _    _   ______   _        _         ____   
| |  | | |  ____| | |      | |       / __ \  
| |__| | | |__    | |      | |      | |  | | 
|  __  | |  __|   | |      | |      | |  | | 
| |  | | | |____  | |____  | |____  | |__| | 
|_|  |_| |______| |______| |______|  \____/  
                                             
                                             
`},
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

	asciiTemplateByte, err := os.ReadFile("./standard.txt")
	if err != nil {
		os.Stdout.WriteString("Error reading file: standard.txt\n")
		os.Exit(1)
	}

	// Split asciiTemplate by double newline ("\n\n") to get individual ASCII characters.
	asciiCharacters := strings.Split(string(asciiTemplateByte), "\n\n")

	// Initialize asciiTable (2D table) (using "make" to avoid out of range).
	asciiTable := make([][]string, len(asciiCharacters))

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
