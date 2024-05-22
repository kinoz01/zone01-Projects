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
			}
		}
	}
	return intColorMap
}

// Find indices of occurence of a substring (Coloring Characters) in a string (userText)
func FindSubstringIndices(colorChars, userText string) []int {
	indices := []int{}
	for i := 0; i < len(userText)-len(colorChars)+1; i++ {
		if userText[i:i+len(colorChars)] == colorChars {
			for j := i; j < i+len(colorChars); j++ {
				indices = append(indices, j) // just the normal index function but we add this loop to get the indice of each character.
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
