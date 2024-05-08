package asciiart

import "regexp"

func PrePrint(userText string) (string, bool) {
	quit := true
	if len(userText) == 0 {
		return "", quit
	}
	if userText == `\n` {
		return "\n", quit
	}
	// Searching for invalid ascii to avoid out of range panic.
	for _, userTextChar := range userText {
		asciiIndex := int(userTextChar)
		if asciiIndex-32 < 0 || asciiIndex-32 >= 95 {
			return "🚨 Found an Invalid Ascii Character.", quit
		}
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	userText = re.ReplaceAllString(userText, "$1")
	quit = false
	return userText, quit
}
