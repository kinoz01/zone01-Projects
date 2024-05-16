package asciiart

import "regexp"

func PrePrint(userText string) (string, bool) {
	
	if len(userText) == 0 {
		return "", true
	}
	if userText == `\n` {
		return "\n", true
	}
	// Searching for invalid ascii to avoid out of range panic.
	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= 95 {
			return "🚨 Found an Invalid Ascii Character.\n", true
		}
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")
	return userText, false
}
