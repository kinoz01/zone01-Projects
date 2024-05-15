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
	reverseErr = "something"
)

var reAlign, reColor, reReverse, reOutput, reFlag *regexp.Regexp

func InitFlagPatterns() {
	reColor = regexp.MustCompile(`\A--color=(\S+)$`)
	reOutput = regexp.MustCompile(`\A--output=(\S+.txt)$`)
	reAlign = regexp.MustCompile(`\A--align=(center|justify|left|right)$`)
	reReverse = regexp.MustCompile(`\A--reverse=(\S+)$`)
	reFlag = regexp.MustCompile(`(--align=|--color=|--output)`)
}

// This function handle all possible arguments errors and quit the program if an error is found.
func ArgsErrors(args []string) (error, bool) {

	if len(args) == 0 {
		return fmt.Errorf("empty input"), true
	}

	argString := strings.Join(args, " ")

	// Quit when a color flag is directly next to a color flag.
	reError := regexp.MustCompile(`--color=(\S+) (--output|--align|--reverse|--color)`)
	if reError.MatchString(argString) {
		err := fmt.Errorf("found a color flag with no value")
		return err, true
	}

	// Quit when user args end with a flag.
	reError2 := regexp.MustCompile(`(--output=\S+$|--align=\S+$|--color=\S+$)`)
	if reError2.MatchString(argString) {
		err := fmt.Errorf("please enter a valid text")
		return err, true
	}

	InitFlagPatterns()

	var a, o, r, c int
	reColorGeneral := regexp.MustCompile(`--color`)
	reOutputGeneral := regexp.MustCompile(`--output`)
	reAlignGeneral := regexp.MustCompile(`--align`)
	reReverseGeneral := regexp.MustCompile(`--reverse`)

	// Ranging over args to find bad flags (flags specification/globalisation).
	for _, arg := range args {
		if reAlignGeneral.MatchString(arg) {
			if reAlign.MatchString(arg) {
				a++
				continue
			} else {
				err := fmt.Errorf("invalid flag: %s", arg)
				return err, true
			}
		}
		if reOutputGeneral.MatchString(arg) {
			if reOutput.MatchString(arg) {
				o++
				continue
			} else {
				err := fmt.Errorf("invalid flag: %s", arg)
				return err, true
			}
		}
		if reReverseGeneral.MatchString(arg) {
			if reReverse.MatchString(arg) {
				reverse = true
				r++
				continue
			} else {
				err := fmt.Errorf("invalid flag: %s", arg)
				return err, true
			}
		}
		if reColorGeneral.MatchString(arg) {
			if reColor.MatchString(arg) {
				if IsValidColor(reColor.FindStringSubmatch(arg)[1]) == "" { // check if colors in color flags are valid.
					err := fmt.Errorf("%s is an invalid color", reColor.FindStringSubmatch(arg)[1])
					return err, true
				}
				c++
				continue
			} else {
				err := fmt.Errorf("invalid flag: %s", arg)
				return err, true
			}
		}
	}
	if reverse && len(args) > 1 {
		err := fmt.Errorf("please enter only one argument to reverse")
		return err, true
	}
	if o > 1 || a > 1 {
		err := fmt.Errorf("too many flags")
		return err, true
	}
	if c == 0 && len(args) > 4 {
		err := fmt.Errorf("too many arguments")
		return err, true
	}

	if len(args) == 2 && GetAsciiTemplateByte(args[len(args)-1]) == nil && !reFlag.MatchString(args[len(args)-2]) {
		err := fmt.Errorf("too many arguments")
		return err, true
	}
	if len(args) >= 3 && GetAsciiTemplateByte(args[len(args)-1]) == nil && !reFlag.MatchString(args[len(args)-2]) && !reColor.MatchString(args[len(args)-3]) {
		err := fmt.Errorf("too many arguments")
		return err, true
	}

	var rmStrings []string
	if GetAsciiTemplateByte(args[len(args)-1]) != nil {
		rmStrings = args[:len(args)-1]
	} else {
		rmStrings = args
	}
	
	// filtring out when strings are allowed (removing strings after a flag exept for color flag and at the end).
	for i, arg := range rmStrings {
		if (reOutput.MatchString(arg) || reAlign.MatchString(arg)) && i+1<len(rmStrings) && i!=len(rmStrings)-2 {
			//fmt.Println("jhhhhh")
			if !reFlag.MatchString(rmStrings[i+1]) {
				err := fmt.Errorf("wrong input: %s", rmStrings[i+1])
				return err, true
			}
		}
		if reColor.MatchString(arg) && i+2<len(rmStrings) && i!=len(rmStrings)-3{
			if !reFlag.MatchString(rmStrings[i+2]) {
				err := fmt.Errorf("wrong input: %s", rmStrings[i+2])
				return err, true
			}
		}
	}

	return nil, false
}
