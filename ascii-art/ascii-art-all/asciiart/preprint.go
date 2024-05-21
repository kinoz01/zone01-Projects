package asciiart

import "regexp"

// Search for invalid ascii to avoid out of range panic, and remove one new line in case of just new lines in userInput.
func PrePrint(userText string) (string, bool) {
	if len(userText) == 0 {
		return "", true
	}
	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= 95 {
			return "🚨 Found an Invalid Ascii Character.\n", true
		}
	}
	re := regexp.MustCompile(`\A((\\n)*)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")
	return userText, false
}
