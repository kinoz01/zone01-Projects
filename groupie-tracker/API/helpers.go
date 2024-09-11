package apiserver

import (
	"strings"
	"unicode"
)

func Search(Word, Type string, Images []string) string {
	// Replace spaces in artistName with hyphens
	Word = strings.ReplaceAll(Word, " ", "-")
	Word = strings.ToLower(Word)
	Word = strings.ReplaceAll(Word, "\"", "")
	Word = strings.ReplaceAll(Word, "'", "")
	Word = strings.ReplaceAll(Word, ".", "")

	// Search for a match in the brandImages slice
	for _, img := range Images {
		if strings.Contains(img, Word) {
			return img
		}
	}
	if Type == "logo" {
		return "https://i.postimg.cc/C1wQ8qwC/no-logo.png"
	} else if Type == "member" {
		return "https://i.postimg.cc/wMTZCsPx/memberplaceholder.jpg"
	} else if Type == "place" {
		return "https://i.postimg.cc/t4K0sJ1D/placeholder-image.webp"
	}
	return "https://i.postimg.cc/wMTZCsPx/memberplaceholder.jpg"
}

func GetMembersImages(Members, Images []string) map[string]string {
	membersMap := make(map[string]string)

	for _, member := range Members {
		memberimage := Search(member, "member", Images)
		membersMap[member] = memberimage
	}
	return membersMap
}

func GetLocationsDates(Dates map[string][]string, Locations, Images []string) map[string][]string {
	LocationsDates := make(map[string][]string)

	for _, place := range Locations {
		placeName := FormatLocation(place)
		LocationsDates[placeName] = append(LocationsDates[placeName], Search(strings.ReplaceAll(place, "_", "-"), "place", Images))
		LocationsDates[placeName] = append(LocationsDates[placeName], Dates[place]...)
	}
	return LocationsDates
}

func GetYoutubeLinks(BandName string, YoutubeLinks map[string][]string) []string {
	return YoutubeLinks[BandName]
}

func FormatLocation(s string) string {

	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return Title(s) // Return original string if it's not in the expected format
	}
	city := Capitalize(parts[0])
	country := Capitalize(parts[1])

	return city + ", " + country
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

// Helper function to capitalize each word
func Capitalize(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = Title(word) // Capitalizes first letter of each word
	}
	return strings.Join(words, " ")
}
