package asciiart

import (
	"fmt"
	"strings"
)

var reverse bool
var ColorAll string

func UserArgs(args []string) (userText, font, alignment, outputFile, reverseInput string, colorMap map[string][]string, quit bool) {
	colorMap = make(map[string][]string)
	font = "standard"
	alignment = "left"
	InitFlagPatterns()

	if err := ArgsErrors(args); err!=nil { // Here I handle all (hopefully) possible input errors.
		// fmt.Println("Error:", err)      // commented this cuz we are restricted to the banal error msg.
		fmt.Println(colorErr)
		return "", "", "", "", "", nil, true
	}
	// if reverse case we need only the ascii input file name to launch reversing mechanism.
	if reverse { 
		return userText, font, alignment, outputFile, strings.TrimPrefix(args[0], "--reverse="), nil, false
	}

	/******* Building upon the fact that the last two words must either contain 'font' and 'user input'
	or just 'user input' as the last word. ***********/
	if len(args) == 1 {
		return args[0], font, alignment, outputFile, reverseInput, nil, false
	} else if len(args) >= 2 {
		if GetAsciiTemplateByte(args[len(args)-1]) != nil { // Ex: --color=red s --align=center shadow || --color=red s sokasoka shadow
			if !reFlag.MatchString(args[len(args)-2]) { // Ex: --color=red s sokasoka shadow
				font = args[len(args)-1]
				userText = args[len(args)-2]
			} else {
				userText = args[len(args)-1] // Ex: --color=red s --align=center shadow
			}
		} else {
			userText = args[len(args)-1] // Ex: --color=red s --align=center Hey || --color=red s sokasoka (already handled all (hopefully) error cases)
		}
	}
	/*********************************************************************************************/
	// here we use submatching to get string of the returns values we will work with. Since we don't have errors we only need to match submatching group with its return value.
	for i, arg := range args {
		if reColorGeneral.MatchString(arg) {
			// We find the Ansi color corresponding to the color string (invalid colors are already handled in args Error, this is just to get the Ansi color value)
			color := IsValidColor(strings.TrimPrefix(arg, "--color="))
			switch i {
			case len(args) - 2: // --color=red hello
				ColorAll = color // coloAll will be used to color all the string output (ascii Art) using "color" and skip coloring parts of the ascii.
				return args[i+1], font, alignment, outputFile, reverseInput, nil, false
			case len(args) - 3: 
				if GetAsciiTemplateByte(args[len(args)-1]) != nil { // --color=red hello shadow
					ColorAll = color
					return args[i+1], args[i+2], alignment, outputFile, reverseInput, nil, false
				} else { // --color=red h hello
					// we used []string map because we can have multiple strings matching a color. (Ex: --color=red o --color=red n "good morning")
					colorMap[color] = append(colorMap[color], args[i+1]) // map[key] = append(map[key], value) (in case of maping to a slice).
					return args[i+2], font, alignment, outputFile, reverseInput, colorMap, false
				}
			default: // --color=red h --color=orange o --align=justify "hello There" o2
				colorMap[color] = append(colorMap[color], args[i+1])
			}
		} else if reAlign.MatchString(arg) {
			alignment = reAlign.FindStringSubmatch(arg)[1]
		} else if reOutput.MatchString(arg) {
			outputFile = reOutput.FindStringSubmatch(arg)[1]

		}
	}

	// If we have a color and output flags at the same time in arguments string we print little msg and continue without coloring (returning nil map).
	if reOutput.MatchString(strings.Join(args, " ")) && reColor.MatchString(strings.Join(args, " ")) {
		fmt.Println("I can't color a txt output file!")
		return userText, font, alignment, outputFile, reverseInput, nil, false
	}

	return userText, font, alignment, outputFile, reverseInput, colorMap, false
}
