package asciiart

import (
	"strings"
)

// we start from color map to get indices of characters we need to color (as a []int -the string part of map will store the colors),
// because working with "char" in the printing loop will lead us to conflict between chars if they are repeated in userText.
func GetColoringIndices(colorMap map[string][]string, userText string) (intColorMap map[string][]int) {
	intColorMap = make(map[string][]int)
	for keyColor, values := range colorMap {
		for _, colorChars := range values {
			if strings.Contains(userText, colorChars) {
				intColorMap[keyColor] = append(intColorMap[keyColor], FindSubstringIndices(colorChars, userText)...) // map[key] = append(map[key], slice...).
			}
		}
	}
	return RemoveDuplicateIndices(intColorMap)
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
func IsColorIndex(indiceColorMap map[string][]int, j int) (string, bool) {
	for color, intSlice := range indiceColorMap {
		for _, indice := range intSlice {
			if indice == j {
				return color, true
			}
		}
	}
	return "", false
}

func RemoveDuplicateIndices(intColorMap map[string][]int) map[string][]int {
	result := make(map[string][]int)
	coloredIndices := make(map[int]bool)

	// Iterate over the map in reverse order
	keys := make([]string, 0, len(intColorMap))
	for key := range intColorMap {
		keys = append(keys, key)
	}
	for i := len(keys) - 1; i >= 0; i-- {
		color := keys[i]
		intSlice := intColorMap[color]
		for j := len(intSlice) - 1; j >= 0; j-- {
			indice := intSlice[j]
			if _, present := coloredIndices[indice]; present {
				continue
			}
			coloredIndices[indice] = true
			// Prepend the index to maintain the original order in the result
			result[color] = append([]int{indice}, result[color]...)
		}
	}
	return result
}
