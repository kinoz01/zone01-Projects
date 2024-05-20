package asciiart

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	colorErr   = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> \"something\""
	outputErr  = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
	alignErr   = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard"
	fontErr    = "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"
	reverseErr = "Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>"
)

var reAlign, reColor, reReverse, reOutput, reFlag *regexp.Regexp

func InitFlagPatterns() {
	reColor = regexp.MustCompile(`\A--color`)
	reOutput = regexp.MustCompile(`\A--output=(.+.txt)$`)
	reAlign = regexp.MustCompile(`\A--align=(center|justify|left|right)$`)
	reReverse = regexp.MustCompile(`\A--reverse=(\S+)$`)
	reFlag = regexp.MustCompile(`(\A--align|\A--color|\A--output)`)
}

// This function handle all possible arguments errors and quit the program if an error is found.
func ArgsErrors(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("empty input")
	}

	argString := strings.Join(args, " ")

	// Quit when a color flag is directly next to a flag (color flag without characters).
	// Ex: --color=red --align=center "zone 01 oujda" o2
	reError := regexp.MustCompile(`--color=(\S+) (--output|--align|--reverse|--color)`)
	if reError.MatchString(argString) {
		err := fmt.Errorf("found a color flag with no value")
		return err
	}

	// Quit when user args end with a flag.
	// Ex: --color=red g --align=center
	// But this will work --color=red g f--align=center (I considered a flag only something that begin with --)
	reError2 := regexp.MustCompile(`(\A--output=.+$|\A--align=.+$|\A--color=.+$)`)	
	if reError2.MatchString(args[len(args)-1]) {
		err := fmt.Errorf("you can't end your input with a flag")
		return err
	}


	InitFlagPatterns()
	// we will remove duplicates with these ints.
	var a, o, c int

	// These patterns will either be flags or error.
	reOutputGeneral := regexp.MustCompile(`\A--output`)
	reAlignGeneral := regexp.MustCompile(`\A--align`)
	reReverseGeneral := regexp.MustCompile(`\A--reverse`)

	// Ranging over args to find bad flags (flags specification/globalisation).
	// first we match any starting flag pattern (wrong or correct) then we return if we find any wrong flag.
	for _, arg := range args {
		if reAlignGeneral.MatchString(arg) {
			if reAlign.MatchString(arg) {
				a++
				continue
			} else {
				err := fmt.Errorf("invalid flag: %s", arg)
				return err
			}
		}
		if reOutputGeneral.MatchString(arg) {
			if reOutput.MatchString(arg) {
				o++
				continue
			} else {
				err := fmt.Errorf("invalid flag: %s", arg)
				return err
			}
		}
		if reReverseGeneral.MatchString(arg) {
			if reReverse.MatchString(arg) {
				reverse = true
				continue
			} else {
				err := fmt.Errorf("invalid flag: %s", arg)
				return err
			}
		}
		if reColor.MatchString(arg) {
			// check if colors in color flags are valid and as consequence check the validity of the flag in general. 
			if IsValidColor(strings.TrimPrefix(arg, "--color=")) != "" { 
				c++
				continue
			} else {
				err := fmt.Errorf("invalid flag: %s", arg)
				return err
			}
		}
	}
	// reverse situations could have only one argument.
	if reverse && len(args) > 1 {
		err := fmt.Errorf("please enter only one argument to reverse")
		return err
	}
	if o > 1 || a > 1 {
		err := fmt.Errorf("too many flags") // only one output/align flag.
		return err
	}
	// if there is no color flag the max args we could have is 4.
	if c == 0 && len(args) > 4 {
		err := fmt.Errorf("too many arguments") 
		return err
	}

	var rmStrings []string
	// GetAsciiTemplateByte reads the font file and return a nil []byte in case of error.
	// Next we remove the last arg if it's a font so it doesn't interfer with our strings/flags removing logic later.
	if GetAsciiTemplateByte(args[len(args)-1]) != nil {
		rmStrings = args[:len(args)-1] 
	} else {
		rmStrings = args
	}

	// filtring out when strings are allowed (removing strings after a flag exept for color flag and at the end).
	for i, arg := range rmStrings {
		// Ex: --align=center lol h hey
		if (reOutput.MatchString(arg) || reAlign.MatchString(arg)) && i+1 < len(rmStrings) && i != len(rmStrings)-2 { 
			if !reFlag.MatchString(rmStrings[i+1]) {
				err := fmt.Errorf("wrong input: %s", rmStrings[i+1])
				return err
			}
		}
		// Ex: --color=red h h y
		if reColor.MatchString(arg) && i+2 < len(rmStrings) && i != len(rmStrings)-3 { 
			if !reFlag.MatchString(rmStrings[i+2]) {
				err := fmt.Errorf("wrong input: %s", rmStrings[i+2])
				return err
			}
		}
		// Ex: hey hey // Ex2: hey --color=red h hey
		if !reFlag.MatchString(rmStrings[0]) && len(rmStrings) > 1 {
			err := fmt.Errorf("wrong input: %s", rmStrings[0])
			return err
		}
	}

	return nil
}
