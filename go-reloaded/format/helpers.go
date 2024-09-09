package format

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// this function searches the index the last occurrence of a substring within a string
func Index(s, toFind string) int {
	runeS := []rune(s)
	runeF := []rune(toFind)
	for i := len(runeS) - len(runeF); i >= 0; i-- {
		if string(runeS[i:i+len(runeF)]) == toFind {
			return i // Last occurrence of substring found at index i
		}
	}
	return -1 // Substring not found
}

func ConvertFromBaseToBase(s string, a, b int) string {
	// Signature: func ParseInt(s string, base int, bitSize int) (i int64, err error)
	num, err := strconv.ParseInt(s, a, 64)
	if err != nil {
		// fmt.Println("Error converting hex to decimal:", err) // no need to print error
		return s // return the input number if non valid in the called base
	}
	// Signature: func FormatInt(i int64, base int) string 
	return strconv.FormatInt(num, b)
}

// this is a helper function checks if the input string contains any letters.
func ContainsLetters(s string) bool {
	for _, char := range s {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

// this function find in []string a word that contains alphabetic character starting from the end of the slice.
func WordFinder(words []string) string {
	endIndex := len(words)-1
	wordFound := ""
	for i := endIndex-1; i >= 0; i--{
		if ContainsLetters(words[i]) {			
			wordFound = words[i]
			break
		}
	}
	return wordFound
}

// this function check if string is valid hex number.
func IsValidHex(hexStr string) bool {
	var dummy int
	n, err := fmt.Sscanf(hexStr, "%x", &dummy)
	return err == nil && n == 1
}

// this function find in []string a valid hex number starting from the end of the slice.
func HexFinder(words []string) string {
	endIndex := len(words)-1
	for i := endIndex-1; i >= 0; i--{
		if IsValidHex(words[i]) {			
			return words[i]
		}
	}
	return ""
}

// IsValidBinary checks if the given string is a valid binary number.
func IsValidBinary(binaryStr string) bool {
	var dummy int
	n, err := fmt.Sscanf(binaryStr, "%b", &dummy)
	return err == nil && n == 1
}

// BinaryFinder finds the first valid binary number in a slice of strings, starting from the end of the slice.
func BinFinder(words []string) string {
	endIndex := len(words) - 1
	for i := endIndex; i >= 0; i-- {
		if IsValidBinary(words[i]) {
			return words[i]
		}
	}
	return ""
}

// find n "valid" words starting from the end of the slice and return them as a slice.
func FindWords(words []string, n int) []string {
	endIndex := len(words) - 3 //remove the number and the flag
	wordsFound := []string{}
	if n <= 0 {
		fmt.Println("🟠 I can't find 0 words before a flag. Please enter a valid number.")
		return wordsFound
	}
	for i := endIndex; i >= 0; i-- {
		if ContainsLetters(words[i]) {
			wordsFound = append(wordsFound, words[i])
			n--
			if n <= 0 {
				break
			}
		}
	}
	if len(wordsFound) == 0 {
		emptyFlag = true
	}
	return wordsFound
}

// search for a word in a string and replace it rune by rune by another word depending on a flag.
func SearchWordAndReplaceIt(s, word, flag string) string {
	runeS := []rune(s)
	runeF := []rune(word)
	// Preprocess the replacement based on the flag
	var replacement []rune
	switch flag {
	case "(up,":
		replacement = []rune(strings.ToUpper(word))
	case "(low,":
		replacement = []rune(strings.ToLower(word))
	}

	for i := len(runeS) - len(runeF); i >= 0; i-- {
		if string(runeS[i:i+len(runeF)]) == word {
			for j := 0; j < len(runeF); j++ {
				runeS[i+j] = replacement[j]

			}
			return string(runeS)
		}
	}
	return s // Substring not found
}

// SearchWordAndReplaceIt searches for a word in a string and replaces it rune by rune by another word depending on a flag.
func SearchWordAndReplaceItCap(s, word string) string {
	runeS := []rune(s)
	runeF := []rune(word)
	replacement := []rune(Title(strings.ToLower(word)))

	for i := len(runeS) - len(runeF); i >= 0; i-- {
		if string(runeS[i:i+len(runeF)]) == word && (i == 0 || !unicode.IsLetter(runeS[i-1])) && (i+len(runeF) == len(runeS) || !unicode.IsLetter(runeS[i+len(runeF)])) {
			for j := 0; j < len(runeF); j++ {
				runeS[i+j] = replacement[j]
			}
			return string(runeS)
		}
	}
	return s // Substring not found
}

// Trim spaces from both the beginning and the end of the line (line by line).
func TrimSpaces(text string) string {
    var result []string
    lines := strings.Split(text, "\n")
    for _, line := range lines {
        trimmedLine := strings.TrimSpace(line)
        result = append(result, trimmedLine)
    }
    return strings.Join(result, "\n")
}

// Title capitalise the first alphabet character found in a word.
func Title(s string) string {
	s = strings.ToLower(s)
	runeS := []rune(s)
	for i, char := range runeS {
		if unicode.IsLetter(char) { // Check if the character is a letter
			runeS[i] = unicode.ToUpper(char) // Convert to uppercase if it is a letter (slices behaves like pointers.)
			break
		}
	}
	return string(runeS)
}


// Determine if the number in the flag is postive or negative counting (++++ and -----) before the number.
func FlagNumPositive(text string, reFlagNegativeNumber *regexp.Regexp) bool {
	matches := reFlagNegativeNumber.FindStringSubmatch(text)

	signs := matches[2] // This captures the sequence of + and - before the number.
	posCount, negCount := 0, 0

	// Count the number of + and - signs.
	for _, char := range signs {
		switch char {
		case '+':
			posCount++
		case '-':
			negCount++
		}
	}
	// Determine the sign based on the counts of + and -.
	// If the number of - signs is odd, the result is negative; otherwise, it's positive.
	return negCount%2 == 0
}
