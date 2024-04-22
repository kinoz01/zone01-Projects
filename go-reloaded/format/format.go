package format

import (
	"regexp"
	"strconv"
	"strings"
	//"fmt"
)

func FormatFlagHelper(s string) string {
	/******************************** This part handles flags without numbers ********************************/
	words := strings.Fields(s)
	endIndex := len(words)-1
		flag1 := words[endIndex]

	if ((flag1 == "(cap)" || flag1 == "(up)" || flag1 == "(low)") && WordFinder(words) == "") || (flag1 == "(hex)" && HexFinder(words) == "") || (flag1 == "(bin)" && BinFinder(words) == "") {
		return s[:len(s)-len(flag1)-1] // in case we don't find any word to modify we just remove the flag from the string to avoid a panic
	}

	switch flag1 {
	case "(cap)":
		return s[:Index(s, WordFinder(words))] + strings.Title(WordFinder(words)) + s[Index(s, WordFinder(words))+len(WordFinder(words)):len(s)-len(flag1)-1]
	case "(up)":
		return s[:Index(s, WordFinder(words))] + strings.ToUpper(WordFinder(words)) + s[Index(s, WordFinder(words))+len(WordFinder(words)):len(s)-len(flag1)-1]
	case "(low)":
		return s[:Index(s, WordFinder(words))] + strings.ToLower(WordFinder(words)) + s[Index(s, WordFinder(words))+len(WordFinder(words)):len(s)-len(flag1)-1]
	case "(hex)":
		return s[:Index(s, HexFinder(words))] + ConvertFromBaseToBase(HexFinder(words), 16, 10) + s[Index(s, HexFinder(words))+len(HexFinder(words)):len(s)-len(flag1)-1]
    case "(bin)":	
		return s[:Index(s, BinFinder(words))] + ConvertFromBaseToBase(BinFinder(words), 2, 10) + s[Index(s, BinFinder(words))+len(BinFinder(words)):len(s)-len(flag1)-1]
	}
	/**********************************************************************************************************/

	/********************************** This part handles flags with numbers **********************************/
	// if we get to this part this means we for sure have "<number>)" in `flag1 := words[endIndex]`
	temp := words[endIndex]
	flag2 := words[endIndex-1]
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
	//fmt.Println(number)
	
	return s
}


// this function handle regular expression in the form (w, low|up|case|hex|bin|cap)
func Flags(text string) string {
	re := regexp.MustCompile(`\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s*(\d+)\))`)
	for re.MatchString(text) {
		pattern := regexp.MustCompile(`(?s)^(.*?)\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))`)
		match := pattern.FindString(text)

		if len(match) < len(text) && (text[len(match)] == ' ' || text[len(match)] == '\n') {
			text = FormatFlagHelper(match) + text[len(match):]
		} else {
			text = FormatFlagHelper(match) + " " + text[len(match):]
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
    re2 := regexp.MustCompile(`([,.!?;:]+)([^,.!?;:\s])`)
    text = re2.ReplaceAllString(text, "$1 $2")
    
    return text
}

func Apostrophe(text string) string {
	// Regular expression to find quoted sections with potential leading and trailing spaces
	re := regexp.MustCompile(`'\s*[^']+?\s*'`)
	return re.ReplaceAllStringFunc(text, func(m string) string {
		// Trim spaces around the matched section
		trimmed := strings.TrimSpace(m)
		// Ensure single quotes are directly next to the inner content
		innerContent := trimmed[1 : len(trimmed)-1] // Remove the outer quotes
		return "'" + strings.TrimSpace(innerContent) + "'"
	})
}

func BasicGrammar(text string) string {
	re := regexp.MustCompile(`(?i)\ba(\s+)([aeiouh]+)`)
    return re.ReplaceAllString(text, "an$1$2")
}

func RemoveTrailingSpaces(text string) string {
	return ""
}

func RemoveTrailingNewLines(text string) string {
	return ""
}
