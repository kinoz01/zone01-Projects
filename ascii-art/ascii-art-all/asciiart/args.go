package asciiart

import (
	"fmt"
	"strings"
)

var (
	reverse  bool
	ColorAll string
	ColorSlice []string
	LenUserText int
)

// This will create a slice that have the same len of the userText to be filled with colors.
func InitColorSlice() {
	ColorSlice = make([]string, LenUserText)
}

// Handle all user args and return strings data and a bool to quit if true.
func UserArgs(args []string) (userText, font, alignment, reverseInput string, outputFiles []string, quit bool) {

	font = "standard"
	alignment = "left"
	InitFlagPatterns()

	// Here I handle all (hopefully) possible input errors.
	if errMsg, err := ArgsErrors(args); err != nil {
		// fmt.Println("Error:", err)
		fmt.Println(errMsg)
		if BadUserFont {
			fmt.Println(fmt.Errorf("only imported 8-lines fonts are supported"))
		}
		return "", "", "", "", nil, true
	}

	// in reverse case we need only the ascii input file name to launch reversing mechanism.
	if reverse {
		return userText, font, alignment, strings.TrimPrefix(args[0], "--reverse="), outputFiles, false
	}

	/******* Building upon the fact that the last two words must either contain 'font' and 'user input'
	or just 'user input' as the last word. ***********/
	// With GetAsciiTemplateByte we First check if the last argument is a font. Ex: --color=red s --align=center shadow || --color=red s sokasoka shadow.
	// Then  we check if the arg before it, is a flag or a string.
	if len(args) == 1 {
		return args[0], font, alignment, reverseInput, outputFiles, false
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
	LenUserText = len(userText)
	InitColorSlice()
	/*********************************************************************************************/

	var o, c, a int
	// In this loop we use submatching to get string of the returns values we will work with. Since we don't have errors we only need to match submatching group with its return value.
	// IsValidColor find the Ansi color corresponding to the color string (invalid colors are already handled in args Error, this is just to get the Ansi color value)
	for i, arg := range args {
		if reColor.MatchString(arg) {
			c++
			color := IsValidColor(strings.TrimPrefix(arg, "--color="))
			GetColorSlice(color, args[i+1], userText)
		} else if reAlign.MatchString(arg) {
			a++
			alignment = reAlign.FindStringSubmatch(arg)[1]
		} else if reOutput.MatchString(arg) {
			o++
			outputFiles = append(outputFiles, reOutput.FindStringSubmatch(arg)[1])
		}
	}

	// If we have a color and output flags at the same time in arguments string we print little msg and continue without coloring (returning nil map).
	if o >= 1 && c >= 1 {
		fmt.Println("You can't color a txt output file!")
		ColorSlice = make([]string, LenUserText)
		return userText, font, alignment, reverseInput, outputFiles, false
	}
	if o>=1 && a>=1 {
		fmt.Println("You can't center a txt output file!")
		alignment = "left"
	}

	return userText, font, alignment, reverseInput, outputFiles, false
}
