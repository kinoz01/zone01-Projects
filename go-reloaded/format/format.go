package format

import (
	"regexp"
	"strconv"
	"strings"
)

// this is a helper function.
func convertFromBaseToBase(s string, a, b int) string {
	// Signature: func ParseInt(s string, base int, bitSize int) (i int64, err error)
	num, err := strconv.ParseInt(s, a, 64)
	if err != nil {
		// fmt.Println("Error converting hex to decimal:", err) // no need to print error
		return s
	}
	// Signature:
	return strconv.FormatInt(num, b)
}

// this function handle regular expression in the form (w, low|up|case|hex|bin|cap)
func Format1(text string) string {

	// Compile the regular expression
	// captures the word before "(flag)"
	re := regexp.MustCompile(`(\b\w+\b)\s+\((low|up|case|hex|bin|cap)\)`)

	// Perform the replacement
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		word := parts[0]
		flag := parts[1]

		switch flag {
		case "(up)":
			return strings.ToUpper(word)
		case "(low)":
			return strings.ToLower(word)
		case "(cap)":
			return strings.Title(word)
		case "(bin)":
			return convertFromBaseToBase(word, 2, 10)
		case "(hex)":
			return convertFromBaseToBase(word, 16, 10)
		}
		return match
	})
}

// this function handle regular expression in the form (words... (low|up|case, <number>))
func Format2(text string) string {
	// capture any text followed by (low, number), (up, number), or (cap, number)
	re := regexp.MustCompile(`((?:\w+\s+)*)(\w+)\s+(\w+)\s+(\w+)\s+\((low|up|cap),\s*(\d+)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		// Number of words to transform
		num, err := strconv.Atoi(match[6])
		if err != nil {
			continue // skip processing this match if the number can't be parsed
		}

		// Total words captured before the control phrase
		allWords := strings.Fields(match[1] + match[2] + " " + match[3] + " " + match[4])
		if num > len(allWords) {
			num = len(allWords) // Prevent index out of range error
		}

		// Find the index to start transformations
		startIndex := len(allWords) - num
		for i := startIndex; i < len(allWords); i++ {
			switch match[5] {
			case "low":
				allWords[i] = strings.ToLower(allWords[i])
			case "up":
				allWords[i] = strings.ToUpper(allWords[i])
			case "cap":
				allWords[i] = strings.Title(allWords[i])
			}
		}

		// Construct the modified segment without the control phrase
		modifiedSegment := strings.Join(allWords, " ")
		fullMatch := match[0]

		// Replace the original segment with the modified segment in the text
		text = strings.Replace(text, fullMatch, modifiedSegment, 1)
	}

	return text
}

func Punctuation(text string) string {
    // Regex to find spaces before punctuation
    re1 := regexp.MustCompile(`\s+([,.!?;:])`)
    text = re1.ReplaceAllString(text, "$1")
    
    // Regex to adjust space after punctuation
    re2 := regexp.MustCompile(`([,.!?;:]+)(\s*)`)
    text = re2.ReplaceAllString(text, "$1 ")
    
    // Regex to trim any excess whitespace after punctuation (if any)
    re3 := regexp.MustCompile(`([,.!?;:]+)\s+`)
    text = re3.ReplaceAllString(text, "$1 ")
    
    // Trim any space at the end of the string if necessary
    re4 := regexp.MustCompile(`\s+$`)
    text = re4.ReplaceAllString(text, "")
    
    return text
}

// func Apostrophe(text string) string {
	
// }
