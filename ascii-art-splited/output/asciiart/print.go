package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func Print(userText, font string) string {
	var output string
	asciiTemplateByte, err := os.ReadFile("./banners/" + font + ".txt")
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return ""
	}
	asciiTemplate := strings.ReplaceAll(string(asciiTemplateByte), "\r", "")
	asciiCharacters := strings.Split(string(asciiTemplate), "\n\n")
	asciiTable := make([][]string, len(asciiCharacters))

	for i := range asciiCharacters {
		lines := strings.Split(asciiCharacters[i], "\n")
		asciiTable[i] = append(asciiTable[i], lines...)
	}

	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" {
			output += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range userLine {
				output += asciiTable[int(char)-32][i]
			}
			output += "\n"
		}
	}
	return output
}
