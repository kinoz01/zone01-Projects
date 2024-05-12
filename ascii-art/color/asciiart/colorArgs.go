package asciiart

import (
	"fmt"
	"regexp"
	"strings"
)

// This function handle all --color flags, it store the color flags values in a map and then remove them from the user args so we can handle the other flags next separately.
// the choice of map[string][]string come out because color flags can have multiple strings colored with the same color.
func GetColorMap(args []string) ([]string, map[string][]string, bool) {
	argsString := strings.Join(args, " ")

	//oldArgs := args
	//var newArgs []string // return user args after removing --color flags.
	var tempArgs []string
	colorMap := make(map[string][]string)

	// return in case of invalid color flag or when a flag is directly next to a color flag or when there is a non flag string between two flags except color flag.
	reError := regexp.MustCompile(`--color=(\S+) (--output|--align|--reverse|--color)`)
	reError2 := regexp.MustCompile(`(--color[^=])|(--color= )|(--color=(\S+)$)`)	
	if reError.MatchString(argsString) || reError2.MatchString(argsString) {
		fmt.Println(colorUsageErr)
		return nil, nil, true
	}

	noLookahead := argsString
	reError3 := regexp.MustCompile(`(--output|--reverse|--align)=\S+ (\S+) (--output|--reverse|--align|--color)=\S+`)
	reFlags := regexp.MustCompile(`(--output=|--reverse=|--align=|--color=)`)
	for reError3.MatchString(noLookahead) {
		checkForFlag := reError3.FindStringSubmatch(noLookahead)
		if !reFlags.MatchString(checkForFlag[2]){
			fmt.Println(colorUsageErr)
			return nil, nil, true
		} else {
			noLookahead = reError3.ReplaceAllString(noLookahead, "")
		}
	}

	// get out and continue if there is no color flag at all.
	reColor := regexp.MustCompile(`--color=(\S+)`)
	if !reColor.MatchString(argsString) {
		return args, nil, false
	}

	// lines 49 ---> 60 handle cases where we have "--color hello something (banner/notBanner)" or "--color hello".
	// Storing userInput and banner.
	if len(args) >= 3 && reColor.MatchString(args[len(args)-3]) {
		if IsBanner(args[len(args)-1]) {
			tempArgs = append(tempArgs, args[len(args)-2], args[len(args)-1])
		} else {
			tempArgs = append(tempArgs, args[len(args)-1]) // in case it's not a banner we don't need to store the element len(args)-2.
		}
		args = args[:len(args)-1] // remove last because in either case it will not be treated by the flag --color.
	} else if len(args) >= 2 && reColor.MatchString(args[len(args)-2]) {
		tempArgs = append(tempArgs, args[len(args)-1])
	}

	argsString = strings.Join(args, " ")

	// creating the color map and removing --color flags along with their correspending characters.
	reColorAndChar := regexp.MustCompile(`--color=(\S+) (\S+)`)
	argsString = reColorAndChar.ReplaceAllStringFunc(argsString, func(match string) string {
		submatches := reColorAndChar.FindStringSubmatch(match)
		color := submatches[1]
		chars := submatches[2]
		if IsValidColor(color) != "" {
			colorMap[IsValidColor(color)] = append(colorMap[IsValidColor(color)], chars) // map[key] = append(map[key], value) (in case of maping to a slice).
		}
		return ""
	})

	args = strings.Fields(argsString)
	args = append(args, tempArgs...)

	reOutput := regexp.MustCompile(`--output`)
	if reOutput.MatchString(argsString) {
		fmt.Println("I can't color a txt output file!")
		return args, nil, false
	}

	// if user enter the same characters for two or more different colors. But in case of overlap this won't work.
	if SameStringForTwoColors(colorMap) {
		fmt.Println("You can't color a text with multiple colors.")
		return args, nil, false
	}

	return args, colorMap, false
}
