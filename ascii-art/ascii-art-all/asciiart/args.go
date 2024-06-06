package asciiart

import (
	"fmt"
	"strings"
)

var (
	reverse  bool
	ColorAll string
	NewLinesColorMaP map[string]map[int][]int
)

func InitNewLinesColorMaP() {
	NewLinesColorMaP = make(map[string]map[int][]int)
}

// Handle all user args and return strings data and a bool to quit if true.
func UserArgs(args []string) (userText, font, alignment, reverseInput string, outputFiles []string, colorMap map[string][]string, quit bool) {
	colorMap = make(map[string][]string)
	font = "standard"
	alignment = "left"
	InitFlagPatterns()
	InitNewLinesColorMaP()

	// Here I handle all (hopefully) possible input errors.
	if errMsg, err := ArgsErrors(args); err != nil {
		// fmt.Println("Error:", err)
		fmt.Println(errMsg)
		if BadUserFont {
			fmt.Println(fmt.Errorf("only imported 8-lines fonts are supported"))
		}
		return "", "", "", "", nil, nil, true
	}

	// in reverse case we need only the ascii input file name to launch reversing mechanism.
	if reverse {
		return userText, font, alignment, strings.TrimPrefix(args[0], "--reverse="), outputFiles, nil, false
	}

	/******* Building upon the fact that the last two words must either contain 'font' and 'user input'
	or just 'user input' as the last word. ***********/
	// With GetAsciiTemplateByte we First check if the last argument is a font. Ex: --color=red s --align=center shadow || --color=red s sokasoka shadow.
	// Then  we check if the arg before it, is a flag or a string.
	if len(args) == 1 {
		return args[0], font, alignment, reverseInput, outputFiles, nil, false
	} else if len(args) >= 2 {
		if GetAsciiTemplateByte(args[len(args)-1]) != nil && !reFlag.MatchString(args[len(args)-2]) {
			// Ex: --color=red s sokasoka shadow
			font = args[len(args)-1]
			userText = args[len(args)-2]
		} else {
			// Ex: --color=red s --align=center shadow || Ex: --color=red s --align=center hey
			userText = args[len(args)-1]
		}
	}
	/*********************************************************************************************/

	var o, c int
	for i, arg := range args {
		if reColor.MatchString(arg) {
			c++
			color := IsValidColor(strings.TrimPrefix(arg, "--color="))
			switch {
			case i == len(args)-2, (i == len(args)-3 && GetAsciiTemplateByte(args[len(args)-1]) != nil):
				ColorAll = color
			default:
				// if strings.Contains(args[i+1], `\n`){
				// 	NewLinesColorMaP[color] = 
				// }
				colorMap[color] = append(colorMap[color], args[i+1])
			}
		} else if reAlign.MatchString(arg) {
			alignment = reAlign.FindStringSubmatch(arg)[1]
		} else if reOutput.MatchString(arg) {
			o++
			outputFiles = append(outputFiles, reOutput.FindStringSubmatch(arg)[1])
		}
	}

	// If we have a color and output flags at the same time in arguments string we print little msg and continue without coloring (returning nil map).
	if o >= 1 && c >= 1 {
		fmt.Println("You can't color a txt output file!")
		ColorAll = ""
		return userText, font, alignment, reverseInput, outputFiles, nil, false
	}

	return userText, font, alignment, reverseInput, outputFiles, colorMap, false
}

/************** Explanation of the loop in this function: *******************/

// In this loop we use submatching to get string of the returns values we will work with. Since we don't have errors we only need to match submatching group with its return value.
// IsValidColor find the Ansi color corresponding to the color string (invalid colors are already handled in args Error, this is just to get the Ansi color value)
// coloAll will be used to color all the string output (ascii Art) using "color" and skip coloring parts of the ascii.
// colorAll turns out to be necessary to handle the case of "newlines (eg, \\n)" in user input. if we remove it we won't get the correct result since we are
// running "strings.Contains" after spliting with "\\n" and now we are appending the whole userText to the ColorMap.
// we used []string map because we can have multiple strings matching a color. (Ex: --color=red o --color=red n "good morning")
// map[key] = append(map[key], value) (in case of maping to a slice).
