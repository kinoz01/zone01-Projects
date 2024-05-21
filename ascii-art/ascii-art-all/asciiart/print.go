package asciiart

import (
	"strings"
)

// construct the final ascii art string.
func PrintAsciiArt(userText, alignement string, asciiTable [][]string, terminalWidth int, colorMap map[string][]string) string {
	var AsciiArt string
	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" {
			AsciiArt += "\n"
			continue
		}
		lenAscii := GetAsciiLineLen(userLine, asciiTable)
		AsciiArt += PrintAsciiLine(userLine, alignement, asciiTable, lenAscii, terminalWidth, colorMap)
	}
	return AsciiArt
}

// construct the user line ascii art using the asciiTable.
func PrintAsciiLine(userLine, alignement string, asciiTable [][]string, lenAscii, terminalWidth int, colorMap map[string][]string) string {
	var output string
	var justify bool
	for i := 0; i < fontLines; i++ {
		switch alignement {
		case "left":
			output += ""
		case "center":
			if terminalWidth-lenAscii > 0 {
				output += strings.Repeat(" ", (terminalWidth-lenAscii)/2)
			}
		case "right":
			if terminalWidth-lenAscii > 0 {
				output += strings.Repeat(" ", terminalWidth-lenAscii)
			}
		case "justify":
			justify = true
			userLine = strings.Join(strings.Fields(userLine), " ")
		}
		for j, char := range userLine {
			if char == ' ' && justify {
				output += GetJustifySpace(terminalWidth, userLine, asciiTable)
				continue
			}
			if color, paint := IsColorIndex(GetColoringIndices(colorMap, userLine), j); paint && ColorAll == "" {
				output += color + asciiTable[int(char-32)][i] + reset
				continue
			}
			output += ColorAll + asciiTable[int(char-32)][i]

		}
		output += "\n"
	}

	return output
}
