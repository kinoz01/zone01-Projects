package format

import (
	"fmt"
	"regexp"
	"strings"
)

func GetUserReact(text string) string {

	reFlagInsideStuff := regexp.MustCompile(`(?i)\([^a-zA-Z]*(low|up|cap|hex|bin)[^\w]*\)`)
	text = reFlagInsideStuff.ReplaceAllStringFunc(text, func(match string) string {
		UserChooseNo := match
		match = strings.ToLower(match)
		switch match {
		case "(cap)", "(bin)", "(hex)", "(up)", "(low)":
			return match
		}

		flags := []string{"cap", "bin", "hex", "up", "low"}
		// Iterate over the flags to check if any flag without a number is present.
		for _, flag := range flags {
			if HasFlagWithoutNum(match, flag) {			
				if GetUserInput(match) == "y" {
					return fmt.Sprintf("(%s)", flag)
				} else {
					// Return the original match if user input is not "y"
					return UserChooseNo
				}
			}
		}
		return match
	})

	reFlagInsideStuff2 := regexp.MustCompile(`(?s)(?i)\([^a-zA-Z]*(low|up|cap)[^a-zA-Z]+\)`)
	text = reFlagInsideStuff2.ReplaceAllStringFunc(text, func(match string) string {
		UserChooseNo := match
		match = strings.ToLower(match)
		reExactFlagWithNum := regexp.MustCompile(`^\((low|up|cap), (\d+)\)$`)
		if reExactFlagWithNum.MatchString(match) {
			return match
		}

		reFlagNegativeNumber := regexp.MustCompile(`(?i)\((cap|low|up),\s+([+-]*)(\d+)\s*\)`) // put(\s*) to handle case of non space after flag.
		if reFlagNegativeNumber.MatchString(match) {
			if FlagNumPositive(match, reFlagNegativeNumber) {
				return reFlagNegativeNumber.ReplaceAllString(match, "($1, $3)")
			} else {
				fmt.Println("Flag takes only postive numbers!! Usage: <(flag, num)>.") 
				if GetUserInput(match) == "y" {
					return reFlagNegativeNumber.ReplaceAllString(match, "($1, $3)")
				} else {
					return UserChooseNo
				}
			}
		}
		
		flags := []string{"cap", "up", "low"}
		// Iterate over the flags to check if any flag without a number is present.
		for _, flag := range flags {
			if found, num := HasFlagWithNum(match, flag); found {			
				if GetUserInput(match) == "y" {
					return fmt.Sprintf("(%s, %v)", flag, num)
				} else {
					// Return the original match if user input is not "y"
					return UserChooseNo
				}
			}
		}
		
		return match
	})
	
	return text
}

/********************************** FIRST FUNCT TO RUN (Try to find every case and interact with user to handle them) *********************************/
func FlagsUserReact(text string) string {
	/************************************** Special Case 1 (apostrophe + space at the start) *****************************/
	if len(text) > 2 && text[0] == '\'' && text[1] == ' ' {
		specialCase1 = true
	}
	if text == "" {
		fmt.Println("It appears that you've provided an empty file.")
		return ""
	}
	var prompt string

	text = GetUserReact(text)
	

	/**************************WITHOUT NUMBER ******* flag pattern incomplete "something(flag" ******************** WITHOUT NUMBER****************/
	reFlagIncomplete := regexp.MustCompile(`(?i)\((cap|low|up|bin|hex)([;:}\{$?!.]|\s+|\n+|$)`)
	prompt = "Found the start of a flag pattern \"(flag ...\". Is this a valid flag? (y/n): "
	if reFlagIncomplete.MatchString(text) && GetUserInputPrompt(prompt) == "y" {
		text = reFlagIncomplete.ReplaceAllString(text, " ($1)$2")
	}


/*****************************************************HANDLE SPACES Before And After The Flag*********************************************************/
	/***************** Flag Between two words with no space **********************/
	reFlagNoBoundSpace := regexp.MustCompile(`(?i)(\S)(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))(\S)`)
	prompt = "Found a flag pattern \"<word>(flag)<word>\". Is this a valid flag? (y/n): "
	if reFlagNoBoundSpace.MatchString(text) && GetUserInputPrompt(prompt) == "y" {
		text = reFlagNoBoundSpace.ReplaceAllString(text, "$1 $2 $6")
	}

	/************ Flag WITH NUMBER is close to the word before it ****************/
	reFlagNoSpaceBefore := regexp.MustCompile(`(?i)(\S)\((low|up|cap),\s+(\d+)\)`)
	prompt = "Found a flag pattern \"<word>(flag, number)\". Is this a valid flag? (y/n): "
	if reFlagNoSpaceBefore.MatchString(text) && GetUserInputPrompt(prompt) == "y" {
		text = reFlagNoSpaceBefore.ReplaceAllString(text, "$1 $2$5")
	}


	/********** When a flag is at the begining *******/
	reFlagSoloStart := regexp.MustCompile(`\A\s*(?i)(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))\s+`)
	prompt = "You can't start your text with a flag. Usage: (words) (flag)."
	if reFlagSoloStart.MatchString(text) {
		fmt.Println(prompt)
		text = reFlagSoloStart.ReplaceAllString(text, "")
	}

	
	/***************************************************************************************************************************************************************/
	/**********************WITH NUMBERS ******** flag with negtaive number "(flag, -/+\d)" ************************ WITH NUMBER****************/

	return text
}

