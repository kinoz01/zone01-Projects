package asciiart

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReverseArt(inputAsciiFile string) string {
	var output string

	inputAsciiByte, err := os.ReadFile("./reverse_ex/" + inputAsciiFile)
	if err != nil {
		inputAsciiByte, err = os.ReadFile(inputAsciiFile)
		if err != nil {
			fmt.Println(reverseErr)
			return ""
		}
	}

	inputAscii := string(inputAsciiByte)
	inputAscii = regexp.MustCompile(`(?m)^\n`).ReplaceAllString(inputAscii, "*\n*\n*\n*\n*\n*\n*\n*\n") // to treat multiple new lines.

	asciiArtLines := make([]string, len(strings.Split(inputAscii, "\n"))/8)
	var j int
	var ascii8Lines string

	for i := 0; i < len(strings.Split(inputAscii, "\n")); i++ {
		ascii8Lines += strings.Split(inputAscii, "\n")[i] + "\n"
		if (i+1)%8 == 0 {
			asciiArtLines[j] = ascii8Lines
			ascii8Lines = ""
			j++
		}
	}
	for i := 0; i < len(asciiArtLines); i++ {
		output += ReverseAsciiArtLine(asciiArtLines[i])
	}

	return output
}

func ReverseAsciiArtLine(inputAsciiLine string) string {
	if inputAsciiLine == "*\n*\n*\n*\n*\n*\n*\n*\n" {
		return "\n"
	}
	AsciiTemplate := strings.Split(string(GetAsciiTemplateByte("standard")), "\n\n")
	AsciiTemplate[len(AsciiTemplate)-1] = regexp.MustCompile(`\n\z`).ReplaceAllString(AsciiTemplate[len(AsciiTemplate)-1], "") // remove the last newline in the tidle
	inputAsciiLines := strings.Split(string(inputAsciiLine), "\n")

	var spaceIndex int
	var AsciiCharacter string
	var result []rune
	var foundSpace bool
	var multipleSpacesCounter int

	for i := 0; i < len(inputAsciiLines[0]); i++ {
		if i+1 < len(inputAsciiLines[0]) && i-1 >= 0 && IsAsciiSpace(inputAsciiLines, i) && !IsAsciiSpace(inputAsciiLines, i-1) {
			for j := 0; j < 8; j++ {
				AsciiCharacter += inputAsciiLines[j][spaceIndex:i+1] + "\n"
			}
			result = append(result, GetNormalCharacter(AsciiTemplate, strings.TrimRight(AsciiCharacter, "\n")))
			AsciiCharacter = ""
			spaceIndex = i + 1
			foundSpace = false
		} else if i+1 == len(inputAsciiLines[0]) && IsAsciiSpace(inputAsciiLines, i) && !IsAsciiSpace(inputAsciiLines, i-1) {
			for j := 0; j < 8; j++ {
				AsciiCharacter += inputAsciiLines[j][spaceIndex:] + "\n"
			}
			result = append(result, GetNormalCharacter(AsciiTemplate, strings.TrimRight(AsciiCharacter, "\n")))
		} else if i+1 < len(inputAsciiLines[0]) && IsAsciiSpace(inputAsciiLines, i) && IsAsciiSpace(inputAsciiLines, i+1) && !foundSpace {
			result = append(result, ' ')
			spaceIndex = i + 6
			foundSpace = true
		}
		if foundSpace {
			multipleSpacesCounter++
		}
		if multipleSpacesCounter%6 == 0 {
			foundSpace = false
		}
	}
	return string(result) + "\n"
}

func IsAsciiSpace(inputAsciiLines []string, j int) bool {
	for i := 0; i < 8; i++ {
		if inputAsciiLines[i][j] != ' ' {
			return false
		}
	}
	return true
}

func GetNormalCharacter(AsciiTemplate []string, AsciiCharacter string) rune {
	for i, AsciiTemplateChar := range AsciiTemplate {
		if AsciiCharacter == AsciiTemplateChar {
			return rune(i + 32)
		}
	}
	return '0'
}
