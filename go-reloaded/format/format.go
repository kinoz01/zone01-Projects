package format

import (
	"regexp"
	"strings"
	//"fmt"
)

func FormatFlagHelper(s string) string {
	/******************************** This parts handles flags without numbers ********************************/
	words := strings.Fields(s)
	endIndex := len(words)-1
		flag1 := words[endIndex]

	if WordFinder(words) == "" || HexFinder(words) == "" {
		return s[:len(s)-len(flag1)-1]
	}

	switch flag1 {
	case "(cap)":
		//fmt.Println("hey")
		return s[:Index(s, WordFinder(words))] + strings.Title(WordFinder(words)) + s[Index(s, WordFinder(words))+len(WordFinder(words)):len(s)-len(flag1)-1]
	case "(up)":
		return s[:Index(s, WordFinder(words))] + strings.ToUpper(WordFinder(words)) + s[Index(s, WordFinder(words))+len(WordFinder(words)):len(s)-len(flag1)-1]
	case "(low)":
		//fmt.Println("hey")
		return s[:Index(s, WordFinder(words))] + strings.ToLower(WordFinder(words)) + s[Index(s, WordFinder(words))+len(WordFinder(words)):len(s)-len(flag1)-1]
	//case "(hex)":
	//	return s[:Index(s, HexFinder(words))] + ConvertFromBaseToBase(HexFinder(words), 16, 10) + s[Index(s, HexFinder(words))+len(HexFinder(words)):len(s)-len(flag1)-1]
  //	case "(bin)":	
	}
	/*********************************************************************************************************/
	return s
}


// this function handle regular expression in the form (w, low|up|case|hex|bin|cap)
func Flags(text string) string {
	re := regexp.MustCompile(`\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s*(\d+)\))`)
	for re.MatchString(text) {
		pattern := regexp.MustCompile(`(?s)^(.*?)\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s*(\d+)\))`)
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
	return ""
}

func BasicGrammar(text string) string {
	return ""
}

func RemoveTrailingSpaces(text string) string {
	return ""
}

func RemoveTrailingNewLines(text string) string {
	return ""
}
