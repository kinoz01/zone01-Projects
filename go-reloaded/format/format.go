package format

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var specialCase1 bool // created to handle the case where an apostrophe is the first letter in text.

/*
func InitialiseSpecialCases() {
	patternSpecialCase2 = `'(\n+)`
}*/

// this function keep finding flags (low|up|case|hex|bin|cap) or (low|up|case|hex|bin|cap, <number>)
func Flags(text string) string {

	re := regexp.MustCompile(`(?i)\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))`)

	for re.MatchString(text) {

		pattern := regexp.MustCompile(`(?i)(?s)^(.*?)\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))`)
		match := pattern.FindString(text)

		text = FormatFlagHelper(match) + text[len(match):]

	}
	return text
}

func FormatFlagHelper(s string) string {

	/******************************** This part handles flags without numbers ********************************/
	words := strings.Fields(s)
	endIndex := len(words) - 1
	flag1 := strings.ToLower(words[endIndex])

	s1 := s[:len(s)-len(flag1)-1]

	if ((flag1 == "(cap)" || flag1 == "(up)" || flag1 == "(low)") && WordFinder(words) == "") || (flag1 == "(hex)" && HexFinder(words) == "") || (flag1 == "(bin)" && BinFinder(words) == "") {
		fmt.Printf("Flag %s need valid expression.\n", flag1)
		return s1 // in case we don't find any word to modify we just remove the flag from the string to avoid infinite loop
	}

	switch flag1 {
	case "(cap)":
		return s[:Index(s1, WordFinder(words))] + Title(WordFinder(words))
	case "(up)":
		return s[:Index(s1, WordFinder(words))] + strings.ToUpper(WordFinder(words))
	case "(low)":
		return s[:Index(s1, WordFinder(words))] + strings.ToLower(WordFinder(words))
	case "(hex)":
		return s[:Index(s1, HexFinder(words))] + ConvertFromBaseToBase(HexFinder(words), 16, 10)
	case "(bin)":
		return s[:Index(s1, BinFinder(words))] + ConvertFromBaseToBase(BinFinder(words), 2, 10)
	}
	/**********************************************************************************************************/

	/********************************** This part handles flags with numbers **********************************/
	// if we get to this part this means we for sure have "<number>)" in `flag1 := words[endIndex]`
	temp := words[endIndex]
	flag2 := strings.ToLower(words[endIndex-1])
	removeFlag := words[endIndex-1] + words[endIndex]

	num, _ := strconv.Atoi(temp[:len(temp)-1]) // remove ")" from the number and convert it to int

	switch flag2 {
	case "(up,":
		for _, str := range FindWords(words, num) {
			s = SearchWordAndReplaceIt(s, str, "(up,")
		}
		return s[:len(s)-len(removeFlag)-2]
	case "(low,":
		for _, str := range FindWords(words, num) {
			s = SearchWordAndReplaceIt(s, str, "(low,")
		}
		return s[:len(s)-len(removeFlag)-2]
	case "(cap,":
		for _, str := range FindWords(words, num) {
			s = SearchWordAndReplaceIt(s, str, "(cap,")
		}
		return s[:len(s)-len(removeFlag)-2]
	}

	return s
}

/********************************** FIRST FUNCT TO RUN (Try to find every case and interact with user to handle them) *********************************/
func FlagsUserReact(text string) string {
	/************************************** Special Case 1 (apostrophe + space at the start) *****************************/
	if len(text) > 2 && text[0] == '\'' && text[1] == ' ' {
		specialCase1 = true
	}
	/*reSpecialCase2 := regexp.MustCompile(patternSpecialCase2)
	if reSpecialCase2.MatchString(text) {
		text = reSpecialCase2.ReplaceAllString(text, "'$1")
	}*/

	/********** When a flag is at the begining *******/
	reFlagSoloStart := regexp.MustCompile(`(?i)^(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))\s+`)
	prompt := "A flag with no valid expression was found. Please enter a valid expression before your flag."
	if reFlagSoloStart.MatchString(text) {
		fmt.Println(prompt)
	}
	text = reFlagSoloStart.ReplaceAllString(text, "")


	/*********** Flag Between two words with no space ****************/
	reFlagNoBoundSpace := regexp.MustCompile(`(?i)(\S)(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))(\S)`)
	prompt = "Found a flag pattern \"<word>(flag)<word>\". Is this a valid flag? (y/n): "
	if reFlagNoBoundSpace.MatchString(text) && GetUserInput(prompt) == "y" {
		text = reFlagNoBoundSpace.ReplaceAllString(text, "$1 $2 $6")
	} 

	/************ Flag is close to the word before it ****************/
	reFlagNoSpaceBefore := regexp.MustCompile(`(?i)(\S)(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))`)
	prompt = "Found a flag pattern \"<word>(flag)\". Is this a valid flag? (y/n): "
	if reFlagNoSpaceBefore.MatchString(text) && GetUserInput(prompt) == "y" {
		text = reFlagNoSpaceBefore.ReplaceAllString(text, "$1 $2")
	}

	/************ Flag is close to the word after it ****************/
	reFlagNoSpaceAfter := regexp.MustCompile(`(?i)(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))(\S)`)
	prompt = "Found a flag pattern \"(flag)<word>\". Is this a valid flag? (y/n): "
	if reFlagNoSpaceAfter.MatchString(text) && GetUserInput(prompt) == "y" {
		text = reFlagNoSpaceAfter.ReplaceAllString(text, "$1 $5")
	}

	/**************WITHOUT NUMBER ******* flag with multiple whitespace on the left OR on the right INSIDE ******** WITHOUT NUMBER**********************/
	reFlagSpaceLeftOrRight := regexp.MustCompile(`(?i)(\(\s+(cap|hex|bin|up|low)\))|(\((cap|hex|bin|up|low)\s+\))`)
	prompt = "Found a flag with white space inside \"(+ flag)/(flag +)\". Is this a valid flag? (y/n): "
	if reFlagSpaceLeftOrRight.MatchString(text) && GetUserInput(prompt) == "y" {
		text = reFlagSpaceLeftOrRight.ReplaceAllString(text, "($4)")
	}

	/**************WITHOUT NUMBER ******* flag with multiple whitespace on the left AND on the right INSIDE ******* WITHOUT NUMBER**********************/
	reFlagSpaceLeftAndRight := regexp.MustCompile(`(?i)(\(\s+(cap|hex|bin|up|low)\s+\))`)
	prompt = "Found a flag with white space inside \"(+ flag +)\". Is this a valid flag? (y/n): "
	if reFlagSpaceLeftAndRight.MatchString(text) && GetUserInput(prompt) == "y" {
		text = reFlagSpaceLeftAndRight.ReplaceAllString(text, "($2)")
	}


	/**************************WITHOUT NUMBER ******* flag pattern incomplete "something(flag" ********************* WITHOUT NUMBER****************/
	reFlagIncomplete := regexp.MustCompile(`(?i)\((cap|low|up|bin|hex)(\s+|\n+|$)`)
	prompt = "Found the start of a flag pattern \"(flag ...\". Is this a valid flag? (y/n): "
	if reFlagIncomplete.MatchString(text) && GetUserInput(prompt) == "y" {
		text = reFlagIncomplete.ReplaceAllString(text, "($1)$2")
	}
	text = FlagReFix(text) // run and Re-fix spaces if it found any. (REFIX)


	/***********WITHOUT NUMBER ******* flag pattern incomplete with punctuation "something(flag[,;:...]" ********* WITHOUT NUMBER****************/
	/*reFlagPoncAfter := regexp.MustCompile(`(?i)\((cap|low|up|bin|hex)(\s+)`)
	prompt = "Found the start of a flag pattern \"<word>(flag\". Is this a valid flag? (y/n): "
	if reFlagPoncAfter.MatchString(text) && GetUserInput(prompt) == "y" {
		text = reFlagPoncAfter.ReplaceAllString(text, "$1 ($2)$3")
	}*/


/***************************************************************************************************************************************************************/
	/**********************WITH NUMBERS ******** flag with negtaive number "(flag, -/+\d)" ************************ WITH NUMBER****************/
	reFlagNegativeNumber := regexp.MustCompile(`(?i)\((cap|low|up),\s+([+-]*)(\d+)\)`) // also handle with(*) cases of multiple spaces after ","
	if reFlagNegativeNumber.MatchString(text) {
		if FlagNumPos(text, reFlagNegativeNumber) {
			text = reFlagNegativeNumber.ReplaceAllString(text, "($1, $3)")
		} else {
			prompt = "Flags take only postive numbers!! Usage: <(flag, num)>. Convert to positive? (y/n): "
			if GetUserInput(prompt) == "y" {
				text = reFlagNegativeNumber.ReplaceAllString(text, "($1, $3)")
			}
		}
	}

	return text
}


func Punctuation(text string) string {
	// Remove spaces before punctuation:
	re1 := regexp.MustCompile(`\s+([,.!?;:])`)
	text = re1.ReplaceAllString(text, "$1")

	// Ensure one space after punctuation:
	// when what come after punctuation is not a punctuation or a whitespace.
	re2 := regexp.MustCompile(`([,.!?;:])([^,.'!?;:\s])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	return text
}

func Apostrophe(text string) string {
	re1 := regexp.MustCompile(`('\s+|\s+')`)
	text = re1.ReplaceAllString(text, " $1 ")

	re2 := regexp.MustCompile(`'([,.!?;:\)\(])`)
	text = re2.ReplaceAllString(text, " ' $1")

	re3 := regexp.MustCompile(`([,.!?;:\)\(])'`)
	text = re3.ReplaceAllString(text, "$1 ' ")

	re4 := regexp.MustCompile(`'\s+`)
	count := 0
	text = re4.ReplaceAllStringFunc(text, func(match string) string {
		if count%2 == 0 {
			count++
			return "'"
		} else {
			count++
			return match
		}
	})
	count = 0
	re5 := regexp.MustCompile(`\s+'`)
	text = re5.ReplaceAllStringFunc(text, func(match string) string {
		if count%2 == 1 {
			count++
			return "'"
		} else {
			count++
			return match
		}
	})
	re6 := regexp.MustCompile(`[ ]+'`)
	text = re6.ReplaceAllString(text, " '")
	re7 := regexp.MustCompile(`'[ ]+`)
	text = re7.ReplaceAllString(text, "' ")
	re8 := regexp.MustCompile(`'\s+\n`)
	text = re8.ReplaceAllString(text, "'\n")

	return strings.TrimRight(text, " \t")
}

func BasicGrammar(text string) string {
	re := regexp.MustCompile(`(?i)\b(a)(\s+)([aeiouh])`)
	return re.ReplaceAllString(text, "${1}n$2$3")
}

func RemoveTrailingSpaces(text string) string {
	re := regexp.MustCompile(`[ \t]+`)
	text = re.ReplaceAllString(text, " ")
	text = TrimSpaces(text)
	return strings.TrimSpace(text)
}

func RemoveTrailingNewLines(text string) string {
	re := regexp.MustCompile(`[\n]+`)
	text = re.ReplaceAllString(text, "\n")
	return strings.TrimSpace(text)
}

/***************************************************** THIS IS THE LAST FUNCTION **********************************************************/
func CleanText(text string) string {
	if specialCase1 {
		text = text[1:]
	}

	reSpaces := regexp.MustCompile(`  +`)
	prompt := "Trailing spaces were detected in your input text, do you want to remove them? (y/n): "
	if reSpaces.MatchString(text) && GetUserInput(prompt) == "y" {
		text = RemoveTrailingSpaces(text)
	}

	reNewlines := regexp.MustCompile(`\n\n+`)
	prompt = "Trailing new lines were detected in your input text, do you want to remove them? (y/n): "
	if reNewlines.MatchString(text) && GetUserInput(prompt) == "y" {
		text = RemoveTrailingNewLines(text)
	}

	reSpaceAtBeginOfNewline := regexp.MustCompile(`\n+ `)
	prompt = "Found space at the beginning of a phrase(s) in your text. Do you want to remove it? (y/n): "
	if (reSpaceAtBeginOfNewline.MatchString(text) || text[0] == ' ') && GetUserInput(prompt) == "y" {
		text = TrimSpaces(text)
	}	

	return text
}
