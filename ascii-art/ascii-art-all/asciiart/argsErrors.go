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
	reColor = regexp.MustCompile(`\A--color=(\S+)$`)
	reOutput = regexp.MustCompile(`\A--output=(\S+.txt)$`)
	reAlign = regexp.MustCompile(`\A--align=(center|justify|left|right)$`)
	reReverse = regexp.MustCompile(`\A--reverse=(\S+)$`)
	reFlag = regexp.MustCompile(`(\A--align=|\A--color=|\A--output)`)
}

// This function handle all possible arguments errors and quit the program if an error is found.
func ArgsErrors(args []string) (error, bool) {

	if len(args) == 0 {
		return fmt.Errorf("empty input"), true
	}

	argString := strings.Join(args, " ")

	// Quit when a color flag is directly next to a color flag.
	// Ex: --color=red --align=center "zone 01 oujda" o2
	reError := regexp.MustCompile(`--color=(\S+) (--output|--align|--reverse|--color)`)
	if reError.MatchString(argString) {
		err := fmt.Errorf("found a color flag with no value")
		return err, true
	}

	// Quit when user args end with a flag.
	// Ex: --color=red g --align=center
	// But this will work --color=red g f--align=center (I considered a flag only something that begin with -- if i remove the space from the pattern it will give error)
	reError2 := regexp.MustCompile(` (--output=\S+$|--align=\S+$|--color=\S+$)`)
	if reError2.MatchString(argString) {
		err := fmt.Errorf("please enter a valid text")
		return err, true
	}

	InitFlagPatterns()
	// we will remove duplicates with these.
	var a, o, r, c int 
	// These patterns will either be flags or error.
	reColorGeneral := regexp.MustCompile(`\A--color`)
	reOutputGeneral := regexp.MustCompile(`\A--output`)
	reAlignGeneral := regexp.MustCompile(`\A--align`)
	reReverseGeneral := regexp.MustCompile(`\A--reverse`)

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
	// reverse situation could have only one argument.
	if reverse && len(args) > 1 {
		err := fmt.Errorf("please enter only one argument to reverse")
		return err, true
	}
	if o > 1 || a > 1 {
		err := fmt.Errorf("too many flags") // only one output/align flag.
		return err, true
	}
	if c == 0 && len(args) > 4 {
		err := fmt.Errorf("too many arguments") // if there is no color flag the max args we could have is 4.
		return err, true
	}

	var rmStrings []string
	if GetAsciiTemplateByte(args[len(args)-1]) != nil { // this function reads the font file and return a nil []byte in case of error.
		rmStrings = args[:len(args)-1] // removing font if exist so it doesn't interfer with our strings/flags logic later.
	} else {
		rmStrings = args
	}
	
	// filtring out when strings are allowed (removing strings after a flag exept for color flag and at the end).
	for i, arg := range rmStrings {
		if (reOutput.MatchString(arg) || reAlign.MatchString(arg)) && i+1<len(rmStrings) && i!=len(rmStrings)-2 { // Ex: --align=center lol h hey
			if !reFlag.MatchString(rmStrings[i+1]) {
				err := fmt.Errorf("wrong input: %s", rmStrings[i+1])
				return err, true
			}
		}
		if reColor.MatchString(arg) && i+2<len(rmStrings) && i!=len(rmStrings)-3{ // Ex: --color=red h h y
			if !reFlag.MatchString(rmStrings[i+2]) {
				err := fmt.Errorf("wrong input: %s", rmStrings[i+2])
				return err, true
			}
		}
		// Ex: hey hey // Ex2: hey --color=red h hey
		if !reFlag.MatchString(rmStrings[0]) && len(rmStrings) >1 {
			err := fmt.Errorf("wrong input: %s", rmStrings[0])
			return err, true
		}
	}

	return nil, false
}
