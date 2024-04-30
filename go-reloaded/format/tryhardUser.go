package format

import (
	"fmt"
	"regexp"
)


func FlagsWrongUsage(text string) string {

	if text == "" {
		fmt.Println("🟠 It appears that you've provided an empty file.")
		return ""
	}

	/********** when flag contains spaces *********/
	reFlagHaveSpaces := regexp.MustCompile(`(?i)\(\s*(cap|low|up|hex|bin)\s*\)`)
	text = reFlagHaveSpaces.ReplaceAllString(text, "($1)")
	reFlagHaveSpaces = regexp.MustCompile(`\(\s*(low|up|cap),\s*(\d+)\s*\)`)
	text = reFlagHaveSpaces.ReplaceAllString(text, "($1, $2)")
	

	/********** when flag is close to the left word *********/
	reFlagToRight := regexp.MustCompile(`(?i)(\S)(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))`)
	text = reFlagToRight.ReplaceAllString(text, "$1 $2")

	/********** When a flag is at the begining *******/
	reFlagSoloStart := regexp.MustCompile(`\A\s*(?i)(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))`)
	if reFlagSoloStart.MatchString(text) {
		emptyFlag = true
		text = reFlagSoloStart.ReplaceAllString(text, "")
	
	}


	// to handle flags with signs. Ask user to parse them or ignore them.
	// Defining a new regexp under the func "ReplaceAllStringFunc" is  a very powerful method
	// as I can now replace submatches of each match using regexp replacing syntax "$1, $2..."
	// and in this way I can even work on submatches using "FindAllStringSubmatch".
	reFlagWithSigns := regexp.MustCompile(`(?i)\(\s*(cap|low|up),\s*([\+\-]+)(\d+)\s*\)`)
	text = reFlagWithSigns.ReplaceAllStringFunc(text, func(match string) string {		
		UserChooseNo := match
		reFlagNegativeNumber := regexp.MustCompile(`(?i)\((cap|low|up), ([\+\-]+)(\d+)\)`) // put(\s*) to handle case of non space after flag.
		if reFlagNegativeNumber.MatchString(match) {
			if FlagNumPositive(match, reFlagNegativeNumber) { 
				if GetUserInput(match) == "y" {
					return reFlagNegativeNumber.ReplaceAllString(match, "($1, $3)")
				} else {
					return UserChooseNo
				}
			} else {
				fmt.Println("🚫 Flag takes only postive numbers!! Usage: \"(word) (flag, number)\".") 
				if GetUserInput(match) == "y" {
					return reFlagNegativeNumber.ReplaceAllString(match, "($1, $3)")
				} else {
					return UserChooseNo
				}
			}
		}		
		return match
	})
	
	return text
}
