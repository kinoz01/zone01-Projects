package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func convertHex(s string) string {
	num, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println("erur")
		return s
	}
	// fmt.Println(num)
	str := strconv.FormatInt(num, 10)
	return str
}

func convertBin(s string) string {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("erur")
		return s
	}
	// fmt.Println(num)
	str := strconv.FormatInt(num, 10)
	return str
}

// this function check if string is valid hex number.
func ValidHex(hexStr string) bool {
	if globalRe.MatchString(hexStr) {
		return false
	}
	var dummy int
	n, err := fmt.Sscanf(hexStr, "%x", &dummy)
	return err == nil && n == 1
}

// IsValidBinary checks if the given string is a valid binary number.
func ValidBin(binaryStr string) bool {
	if globalRe.MatchString(binaryStr) {
		return false
	}
	var dummy int
	n, err := fmt.Sscanf(binaryStr, "%b", &dummy)
	return err == nil && n == 1
}

// this is a helper function checks if the input string contains any letters.
func ValidWord(s string) bool {
	if globalRe.MatchString(s) {
		return false
	}
	for _, char := range s {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
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
