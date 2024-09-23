package apiserver

import (
	"groupie/server"
	"strings"
	"unicode"
)

// search for a string inside a list of links to respond with the correct link.
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
		return server.APILinks.OtherLinks[0]
	} else if Type == "member" {
		return server.APILinks.OtherLinks[1]
	} else if Type == "place" {
		return server.APILinks.OtherLinks[2]
	}
	return server.APILinks.OtherLinks[1]
}

// Return a map where the key is a member name and the value is the link to its corresponding image.
func GetMembersImages(Members, Images []string) map[string]string {
	membersMap := make(map[string]string)

	for _, member := range Members {
		memberimage := Search(member, "member", Images)
		membersMap[member] = memberimage
	}
	return membersMap
}

// Return a map where key is locations and values are a slice of string containing dates.
func GetLocationsDates(Dates map[string][]string, Locations []string) map[string][]string {
	LocationsDates := make(map[string][]string)

	for _, place := range Locations {
		placeName := FormatLocation(place)
		LocationsDates[placeName] = append(LocationsDates[placeName], GetApiImage(placeName))
		LocationsDates[placeName] = append(LocationsDates[placeName], Dates[place]...)
	}
	return LocationsDates
}

// Return youtube links depending on artist name.
func GetYoutubeLinks(BandName string, YoutubeLinks map[string][]string) []string {
	return YoutubeLinks[strings.ToLower(BandName)]
}

// Format the location string from reo_dejanero-brazil-------------> Reo Dejaniro, Brazil
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
