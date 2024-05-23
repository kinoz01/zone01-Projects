package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func ReverseArt(inputAsciiFile string) string {
	inputAsciiByte, err := os.ReadFile("./reverse_ex/" + inputAsciiFile)
	if err != nil {
		inputAsciiByte, err = os.ReadFile(inputAsciiFile)
		if err != nil {
			fmt.Println(colorErr)
			return ""
		}
	}

	AsciiTemplate := strings.Split(string(GetAsciiTemplateByte("standard")), "\n\n")
	inputAsciiLines := strings.Split(string(inputAsciiByte), "\n")

	var spaceIndex int
	var AsciiCharacter string
	var result []rune
	var foundSpace bool
	var multiplespacesCounter int

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
			multiplespacesCounter++
		}
		if multiplespacesCounter%6 == 0 {
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
