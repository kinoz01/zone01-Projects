package asciiart

var fontLines int

// this function initialize the number of lines present in the ascii character so we can range on them later.
func InitFontLines(font string) {
	switch font {
	case "graceful":
		fontLines = 4
	case "small":
		fontLines = 5
	case "phoenix", "o2", "starwar", "stop", "varsity":
		fontLines = 7
	case "standard", "shadow", "thinkertoy", "arob", "zigzag", "henry3D", "doom", "tiles", "jacky", "catwalk", "coins":
		fontLines = 8
	case "fire":
		fontLines = 9
	case "jazmine", "matrix":
		fontLines = 10
	case "blocks", "univers":
		fontLines = 11
	case "impossible":
		fontLines = 12
	case "georgi":
		fontLines = 16
	default:
		fontLines = 8
	}
}
