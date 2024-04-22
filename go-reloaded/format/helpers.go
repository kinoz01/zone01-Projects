package format

import (
	"strconv"
	"strings"
	"fmt"
	"unicode"
)

// this function searches the index the last occurrence of a substring within a string
func Index(s string, toFind string) int {
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

// find n words starting from the end of the slice and return them as a slice.
func FindWords(words []string, n int) []string {
	endIndex := len(words) - 3
	wordsFound := []string{}
	for i := endIndex; i >= 0; i-- {
		if ContainsLetters(words[i]) {
			wordsFound = append(wordsFound, words[i])
			n--
			if n <= 0 {
				break
			}
		}
	}
	return wordsFound
}

func SearchWordAndReplaceIt(s, word, flag string) string {
	runeS := []rune(s)
	runeF := []rune(word)
	for i := len(runeS) - len(runeF); i >= 0; i-- {
		if string(runeS[i:i+len(runeF)]) == word {
			for j := 0; j < len(runeF); j++ {
				switch flag {
				case "(up,":
					runeS[i+j] = []rune(strings.ToUpper(word))[j]
				case "(low,":
					runeS[i+j] = []rune(strings.ToLower(word))[j]
				case "(cap,":
					runeS[i+j] = []rune(strings.Title(word))[j]
				}				
			}
			return string(runeS)
		}
	}
	return s // Substring not found
}
