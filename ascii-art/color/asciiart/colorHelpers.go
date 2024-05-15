package asciiart

import (
	"strings"
)

// we start from color map to get indices of characters we need to color (as a []int -the string part of map will store the colors), 
// because working with "char" in the printing loop will lead us to conflict between chars if they are repeated in userText.
func GetColoringIndices(colorMap map[string][]string, userText string) map[string][]int {
	intColorMap := make(map[string][]int)

	for keyColor, values := range colorMap {
		for _, colorChars := range values {
			if strings.Contains(userText, colorChars) {
				intColorMap[keyColor] = append(intColorMap[keyColor], FindSubstringIndices(colorChars, userText)...) // map[key] = append(map[key], slice...).
			} else if strings.Contains(colorChars, userText)  { 
				// sometimes the userText is shorter than colorChars, in this case i want to color all userText but only if it exist in colorChars (colorChars contains userText).
				intColorMap[keyColor] = append(intColorMap[keyColor], FindSubstringIndices(userText, colorChars)...)
			}
		}
	}
	return intColorMap
}

// Find indices of occurence of a substring (generally -but not always- colorChars) in a string (userText)
func FindSubstringIndices(colorChars, userText string) []int { 
	indices := []int{}
	for i := 0; i < len(userText)-len(colorChars)+1; i++ {
		if userText[i:i+len(colorChars)] == colorChars {
			for j := i; j < i+len(colorChars); j++ {
				indices = append(indices, j)
			}
		}
	}
	return indices
}

// check if an indice corresponding to a char in printing loop is present in the intColorMap, 
// if found return the color corresponding to that indice (represented by the key in the map)
func IsColorIndex(indexColorMap map[string][]int, j int) (string, bool) {
	for color, intSlice := range indexColorMap {
		for _, index := range intSlice {
			if index == j {
				return color, true
			}
		}
	}
	return "", false
}

// This function check if user entered same exact character(s)/values for two colors/map keys.
// if true used to return an empty map in GetColorMap func and continue printing without colors.
func SameStringForTwoColors(colorMap map[string][]string) bool {
	seen := []string{}

	for _, values := range colorMap {
		for _, value := range values {
			for _, seenStr := range seen {
				if seenStr == value {
					return true
				}
			}
			seen = append(seen, value)
		}
	}
	return false
}
