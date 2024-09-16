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
	if ascii8Lines != "\n" {
		fmt.Println("your ascii art is not correctly formatted")
		os.Exit(1)
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
		if i-1 >= 0 && IsAsciiSpace(inputAsciiLines, i) && !IsAsciiSpace(inputAsciiLines, i-1) {
			for j := 0; j < 8; j++ {
				AsciiCharacter += inputAsciiLines[j][spaceIndex:i+1] + "\n"
			}
			r, err := GetNormalCharacter(AsciiTemplate, strings.TrimRight(AsciiCharacter, "\n"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			result = append(result, r)		
			AsciiCharacter = ""
			spaceIndex = i + 1
			foundSpace = false
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
	
	CheckLast(AsciiCharacter, AsciiTemplate, inputAsciiLines, spaceIndex)

	return string(result) + "\n"
}

func CheckLast(AsciiCharacter string, AsciiTemplate, inputAsciiLines []string, spaceIndex int) {
	for j := 0; j < 8; j++ {
		AsciiCharacter += inputAsciiLines[j][spaceIndex:] + "\n"
	}
	_, err := GetNormalCharacter(AsciiTemplate, strings.TrimRight(AsciiCharacter, "\n"))
	if err != nil && AsciiCharacter != "\n\n\n\n\n\n\n\n"{
		fmt.Println(err)
		os.Exit(1)
	}

}

func IsAsciiSpace(inputAsciiLines []string, j int) bool {
	for i := 0; i < 8; i++ {
		if j >= len(inputAsciiLines[i]) {
			fmt.Println("your ascii art is not correctly formatted")
			os.Exit(1)
		}
		if inputAsciiLines[i][j] != ' ' {
			return false
		}
	}
	return true
}

func GetNormalCharacter(AsciiTemplate []string, AsciiCharacter string) (rune, error) {
	for i, AsciiTemplateChar := range AsciiTemplate {
		if AsciiCharacter == AsciiTemplateChar {
			return rune(i + 32), nil
		}
	}
	return '0', fmt.Errorf("your ascii art is not correctly formatted")
}
