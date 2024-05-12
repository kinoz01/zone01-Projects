package asciiart

import (
	"strings"
)

func GetColoringIndex(colorMap map[string][]string, userText string) map[string][]int {
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
