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
	endIndex := len(words) - 1
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

	return s
}

// this function handle flags (low|up|case|hex|bin|cap) or (low|up|case|hex|bin|cap, <number>)
func Flags(text string) string {
	re := regexp.MustCompile(`\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s*(\d+)\))`)
	for re.MatchString(text) {
		pattern := regexp.MustCompile(`(?s)^(.*?)\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap),\s+(\d+)\))`)
		match := pattern.FindString(text)

		if len(match) < len(text) && (text[len(match)] == ' ' || text[len(match)] == '\n') {
			text = FormatFlagHelper(match) + text[len(match):] // no need for space if I already have space
		} else {
			text = FormatFlagHelper(match) + " " + text[len(match):] // to add a space in case I have a word sticking to the right of the flag
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
		//fmt.Println(count)
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
	return strings.TrimSpace(text)
}

func RemoveTrailingNewLines(text string) string {
	re := regexp.MustCompile(`[\n]+`)
	text = re.ReplaceAllString(text, "\n")
	return strings.TrimSpace(text)
}
