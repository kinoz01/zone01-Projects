package format

import (
	"regexp"
	"strconv"
	"strings"
	//"fmt"
)

// this is a helper function.
func convertFromBaseToBase(s string, a, b int) string {
	// Signature: func ParseInt(s string, base int, bitSize int) (i int64, err error)
		// int64 specify that I want int64 even in a no 64-bits architecture computer
		// if I do "var i int64" in a 32-bits system I wont get overflow error unlike using just int
	num, err := strconv.ParseInt(s, a, 64)
	if err != nil {
		// fmt.Println("Error converting hex to decimal:", err) // no need to print error
		return s // return the input number if non valid in the called base
	}
	// Signature: func FormatInt(i int64, base int) string 
	return strconv.FormatInt(num, b)
}

// This is a helper function fixing if there is a flag in the end of the text
func FixWhenFlagLast(text string) string{
	if text[len(text)-1] != ' ' || text[len(text)-1] != '\n' {
		return text + " "
	}
	return text
}

// this function handle regular expression in the form (w, low|up|case|hex|bin|cap)
func Format1(text string) string {
	// Compile the regular expression
	re := regexp.MustCompile(`\b(\w+)\b[.,:;']*\s+\((low|up|case|hex|bin|cap)\)\W`) 
	// note that here we included whitespaces lastely so we need to get them back later when replacing

	// Perform the replacement: which means run `repl func` for each match
	// signature: func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

	return re.ReplaceAllStringFunc(text, func(match string) string {
		
		parts := strings.Fields(match)
		word := parts[0]
		flag := parts[1]

		var r string
		switch flag {
		case "(up)":
			r = strings.ToUpper(word) + " "
		case "(low)":
			r = strings.ToLower(word) + " "
		case "(cap)":
			r = strings.Title(word) + " "
		case "(bin)":
			r = convertFromBaseToBase(word, 2, 10) + " "
		case "(hex)":
			r = convertFromBaseToBase(word, 16, 10) + " "
		}

		// Preserve newline characters if present in the original match
		if strings.ContainsAny(match, "\n") {
			r += "\n"
		}

		return r
	})
}


func Format2(text string) string {
	
	flagNumPattern := "\\b(\\w+)\\b([.,:;'\\!\\?]*)\\s+\\((low|up|cap),\\s*(\\d+)\\)\\s"
	re := regexp.MustCompile(flagNumPattern)
	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		num, _ := strconv.Atoi(match[4])
		pattern := ""
		for num > 1 {
			pattern += "\\b(\\w+)\\b([.,:;'\\!\\?]*)\\s+"
			num--
		}
		pattern += flagNumPattern

		rex := regexp.MustCompile(pattern)

		text = rex.ReplaceAllStringFunc(text, func(s string) string {
			parts := strings.Fields(s)
			result := ""
			// Apply transformation to each word except the last two (directive and number)
			for _, word := range parts[:len(parts)-2] {
				switch parts[len(parts)-2] {
				case "(cap,":
					word = strings.Title(word)
				case "(low,":
					word = strings.ToLower(word)
				case "(up,":
					word = strings.ToUpper(word)
				}
				result += word + " " // add the last space character that we overlaped while defining the pattern
				
			}
			if strings.ContainsAny(s, "\n") {
				result += "\n"
			}
			return result 
		})
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
