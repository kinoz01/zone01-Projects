package asciiart

import (
	"strings"
)

// construct the final ascii art string.
func PrintAsciiArt(userText, alignement string, asciiTable [][]string, terminalWidth int) string {
	var AsciiArt string
	var ColorIndex int // Instead of using a global variable I can use a pointer.
	for _, userLine := range strings.Split(userText, `\n`) {
		if userLine == "" {
			AsciiArt += "\n"
			ColorIndex += 2
			continue
		}
		lenAscii := GetAsciiLineLen(userLine, asciiTable)
		AsciiArt += PrintAsciiLine(userLine, alignement, asciiTable, lenAscii, terminalWidth, &ColorIndex)
		ColorIndex += 2 // to skip '\n'
	}
	return AsciiArt
}

// construct the user line (splited by \n) to ascii art using the asciiTable.
func PrintAsciiLine(userLine, alignement string, asciiTable [][]string, lenAscii, terminalWidth int, ColorIndex *int) string {

	temp := *ColorIndex

	var output string
	var justify bool
	for i := 0; i < fontLines; i++ {
		*ColorIndex = temp
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
		for _, char := range userLine {
			if char == ' ' && justify {
				output += GetJustifySpace(terminalWidth, userLine, asciiTable)
				continue
			}
			if ColorSliceEmpty() {
				output += asciiTable[int(char-32)][i]
			} else {
				output += ColorSlice[*ColorIndex] + asciiTable[int(char-32)][i] + reset
				*ColorIndex++
			}
		}
		output += "\n"
	}
	return output
}
