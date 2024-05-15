package asciiart

import (
	"fmt"
	"strings"
)

var reverse bool
var asciiColor string

func UserArgs(args []string) (userText, font, alignment, outputFile, reverseInput string, colorMap map[string][]string, quit bool) {
	colorMap = make(map[string][]string)
	font = "standard"
	alignment = "left"
	InitFlagPatterns()

	if _, quit := ArgsErrors(args); quit { // Here I handle all possible input errors.
		// fmt.Println("Error: ", err)     // we are restricted to the banal error msg.
		fmt.Println(colorErr)
		return "", "", "", "", "", nil, true
	}
	if reverse {
		return userText, font, alignment, outputFile, strings.TrimPrefix(args[0], "--reverse="), nil, false
	}

	/******* Building upon the fact that the last two words must either contain 'font' and 'user input'
	or just 'user input' as the last word. ***********/
	if len(args) == 1 {
		return args[0], font, alignment, outputFile, reverseInput, nil, false
	} else if len(args) >= 2 {
		if GetAsciiTemplateByte(args[len(args)-1]) != nil {
			if !reFlag.MatchString(args[len(args)-2]) {
				font = args[len(args)-1]
				userText = args[len(args)-2]
			} else {
				userText = args[len(args)-1]
			}
		} else {
			userText = args[len(args)-1]
		}
	}
	/*********************************************************************************************/

	for i, arg := range args {
		if reColor.MatchString(arg) {
			color := IsValidColor(reColor.FindStringSubmatch(arg)[1])
			switch i {
			case len(args) - 2:
				asciiColor = color
				return args[i+1], font, alignment, outputFile, reverseInput, nil, false
			case len(args) - 3:
				if GetAsciiTemplateByte(args[len(args)-1]) != nil {
					asciiColor = color
					return args[i+1], args[i+2], alignment, outputFile, reverseInput, nil, false
				} else {
					colorMap[color] = append(colorMap[color], args[i+1]) // map[key] = append(map[key], value) (in case of maping to a slice).
					return args[i+2], font, alignment, outputFile, reverseInput, colorMap, false
				}
			default:
				colorMap[color] = append(colorMap[color], args[i+1])
			}
		} else if reAlign.MatchString(arg) {
			alignment = reAlign.FindStringSubmatch(arg)[1]
		} else if reOutput.MatchString(arg) {
			outputFile = reOutput.FindStringSubmatch(arg)[1]

		} else if reReverse.MatchString(arg) {
			reverseInput = reReverse.FindStringSubmatch(arg)[1]
		}
	}

	if reOutput.MatchString(strings.Join(args, " ")) && reColor.MatchString(strings.Join(args, " ")) {
		fmt.Println("I can't color a txt output file!")
		return userText, font, alignment, outputFile, reverseInput, nil, false
	}

	return userText, font, alignment, outputFile, reverseInput, colorMap, false
}
