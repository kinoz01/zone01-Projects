package asciiart

import (
	"strings"
)

func PrintAsciiArt(userText, alignement string, asciiTable [][]string, terminalWidth int) string {
	var AsciiArt string
	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" {
			AsciiArt += "\n"
			continue
		}
		lenAscii := GetAsciiLineLen(userLine, asciiTable)
		AsciiArt += PrintAsciiLine(userLine, alignement, asciiTable, lenAscii, terminalWidth)
	}
	return AsciiArt
}

func PrintAsciiLine(userLine, alignement string, asciiTable [][]string, lenAscii, terminalWidth int) string {
	var output string
	var row string
	var justify bool
	for i := 0; i < fontLines; i++ {
		switch alignement {
		case "left":
			row = ""
		case "center":
			row = GetCenterSpaces(terminalWidth, lenAscii)
		case "right":
			row = GetRightSpaces(terminalWidth, lenAscii)
		case "justify":
			justify = true
		}
		for _, char := range userLine {
			if char == ' ' && justify {
				row += GetJustifySpace(terminalWidth, userLine, asciiTable)
				continue
			}
			row += asciiTable[int(char-32)][i]
		}
		row += "\n"
		output += row
	}
	return output
}
