package asciiart

var fontLines int

// this function initialize the number of lines present in the ascii character so we can range on them later.
func InitFontLines(font string) {
	switch font {
	case "small":
		fontLines = 5
	case "phoenix", "o2", "starwar", "stop":
		fontLines = 7
	case "standard", "shadow", "thinkertoy", "arob", "zigzag", "henry3D", "doom":
		fontLines = 8
	case "blocks":
		fontLines = 11
	default:
		fontLines = 8
	}
}
