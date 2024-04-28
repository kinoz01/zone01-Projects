package format

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var emptyFlag bool

// this function keep finding flags (low|up|case|hex|bin|cap) or (low|up|case|hex|bin|cap, <number>)
/**************************************************** This part handles flags WITHOUT NUMBES ***********************************************************/
// this function handle flags (low|up|case|hex|bin|cap).
func Flags(text string) string {
	reNoNumFlag := regexp.MustCompile(`(?i)(((\s|\n)\((cap|low|up|hex|bin)\)(\s|$)))|((\s|\n)\((cap|low|up), \d+\))`)
	for reNoNumFlag.MatchString(text) {

		pattern := regexp.MustCompile(`(?i)(?s)(.*?)\s+(\((cap|low|up|hex|bin)\)|\((low|up|cap), (\d+)\))`)
		match := pattern.FindString(text)

		if len(match) < len(text) && (text[len(match)] != ' ' && text[len(match)] != '\n') {
			text = FormatFlags(match) + " " + text[len(match):] // to add a space in case I have a word sticking to the right of the flag
		} else {
			text = FormatFlags(match) + text[len(match):] // no need for space if I already have space
		}
	}

	return text
}

func FormatFlags(s string) string {

	words := strings.Fields(s)
	flag := strings.ToLower(words[len(words)-1])

	s1 := s[:len(s)-len(flag)-1]

	if ((flag == "(cap)" || flag == "(up)" || flag == "(low)") && WordFinder(words) == "") || (flag == "(hex)" && HexFinder(words) == "") || (flag == "(bin)" && BinFinder(words) == "") {
		emptyFlag = true
		return s1 // in case we don't find any word to modify we just remove the flag from the string to avoid infinite loop
	}

	switch flag {
	case "(cap)":
		return s1[:Index(s1, WordFinder(words))] + Title(WordFinder(words)) + s1[Index(s1, WordFinder(words))+len(WordFinder(words)):]
	case "(up)":
		return s1[:Index(s1, WordFinder(words))] + strings.ToUpper(WordFinder(words)) + s1[Index(s1, WordFinder(words))+len(WordFinder(words)):]
	case "(low)":
		return s1[:Index(s1, WordFinder(words))] + strings.ToLower(WordFinder(words)) + s1[Index(s1, WordFinder(words))+len(WordFinder(words)):]
	case "(hex)":
		return s1[:Index(s1, HexFinder(words))] + ConvertFromBaseToBase(HexFinder(words), 16, 10) + s1[Index(s1, HexFinder(words))+len(HexFinder(words)):]
	case "(bin)":
		return s1[:Index(s1, BinFinder(words))] + ConvertFromBaseToBase(BinFinder(words), 2, 10) + s1[Index(s1, BinFinder(words))+len(BinFinder(words)):]
	}
	/**********************************************************************************************************/

	/********************************** This part handles flags with numbers **********************************/
	// if we get to this part this means we for sure have "<number>)" in `flag1 := words[endIndex]`
	temp := words[len(words)-1]
	flag2 := strings.ToLower(words[len(words)-2]) // we have a flag on the form (up, 2) so we the index of the flag is in len(words)-2 .
	if temp[len(temp)-1] != ')' {
		return s
	}
	removeFlag := words[len(words)-2] + words[len(words)-1] // tha flag is equal to: "(flag," + "num)"

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

func Punctuation(text string) string {
	// Remove spaces before punctuation:
	re1 := regexp.MustCompile(` +([,.!?;:])`)
	text = re1.ReplaceAllString(text, "$1")

	// Ensure one space after punctuation:
	// when what come after punctuation is not a punctuation or a whitespace.
	re2 := regexp.MustCompile(`([,.!?;:])([^,.'!?;:\s])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	return text
}

func Apostrophe(text string) string {

	lines := strings.Split(text, "\n")
	newLines := []string{}
	for _, line := range lines {
		re1 := regexp.MustCompile(`('\s+|\s+')`)
		line = re1.ReplaceAllString(line, " $1 ")
		re12 := regexp.MustCompile(`\A'`)
		line = re12.ReplaceAllString(line, " ' ")

		re2 := regexp.MustCompile(`'([\[\]{},.!'?;:)^\(\n])`)
		line = re2.ReplaceAllString(line, " ' $1")

		re3 := regexp.MustCompile(`([,.!?;:\)'(\n])'`)
		line = re3.ReplaceAllString(line, "$1 ' ")

		re4 := regexp.MustCompile(`'\s+`)
		count := 0
		line = re4.ReplaceAllStringFunc(line, func(match string) string {
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
		line = re5.ReplaceAllStringFunc(line, func(match string) string {
			if count%2 == 1 {
				count++
				return "'"
			} else {
				count++
				return match
			}
		})
		newLines = append(newLines, line)
	}

	text = strings.Join(newLines, "\n")

	re6 := regexp.MustCompile(`[ ]+'`)
	text = re6.ReplaceAllString(text, " '")
	re7 := regexp.MustCompile(`'[ ]+`)
	text = re7.ReplaceAllString(text, "' ")
	re8 := regexp.MustCompile(`' +(\n+) +`)
	text = re8.ReplaceAllString(text, "'$1")
	re9 := regexp.MustCompile(`\n '`)
	text = re9.ReplaceAllString(text, "\n'")
	re10 := regexp.MustCompile(`' \n`)
	text = re10.ReplaceAllString(text, "'\n")
	re11 := regexp.MustCompile(`\A '`)
	text = re11.ReplaceAllString(text, "'")

	return strings.TrimRight(text, " \t")
}

func BasicGrammar(text string) string {
	re := regexp.MustCompile(`(?i)([^âéèĥ]\ba)(\s+)([aeiouh])`)
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

	reSpaces := regexp.MustCompile(`  +`)
	prompt := "🤷 Trailing spaces were detected in your input text, do you want to remove them? (y/n): "
	if reSpaces.MatchString(text) && GetUserInputPrompt(prompt) == "y" {
		text = RemoveTrailingSpaces(text)
	}

	reNewlines := regexp.MustCompile(`\n\n+`)
	prompt = "🤷 Trailing new lines were detected in your input text, do you want to remove them? (y/n): "
	if reNewlines.MatchString(text) && GetUserInputPrompt(prompt) == "y" {
		text = RemoveTrailingNewLines(text)
	}

	reSpaceAtBeginOfNewline := regexp.MustCompile(`\n+ `)
	reSpaceAtbeginOfText := regexp.MustCompile(`\A +`)
	prompt = "🤷 Found space at the beginning of a phrase(s) in your text. Do you want to remove it? (y/n): "
	if len(text) > 1 && (reSpaceAtBeginOfNewline.MatchString(text) || reSpaceAtbeginOfText.MatchString(text)) && GetUserInputPrompt(prompt) == "y" {
		text = TrimSpaces(text)
	}

	if emptyFlag {
		fmt.Println("🟠 Invalid flags detected. A flag should be called after a valid expression. All flags will be removed.")
	}

	return text
}
