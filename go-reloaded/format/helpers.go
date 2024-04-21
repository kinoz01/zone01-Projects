package format

import (
	"strconv"
	//"strings"
	"unicode"
	"fmt"
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

// this is a helper function to convert from base to base.
func ConvertFromBaseToBase(s string, a, b int) string {
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

// this is a helper function checks if the input string contains any letter.
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

func isValidHex(hexStr string) bool {
	var dummy int
	n, err := fmt.Sscanf(hexStr, "%x", &dummy)
	return err == nil && n == 1
}

func HexFinder(words []string) string {
	endIndex := len(words)-1
	hexFound := ""
	for i := endIndex-1; i >= 0; i--{
		if IsHexNumber(words[i]) {			
			hexFound = words[i]
			break
		}
	}
	return hexFound
}