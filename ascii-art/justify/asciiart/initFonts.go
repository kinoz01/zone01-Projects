package asciiart

var fontLines int

func InitFontLines(font string) {
	switch font {
	case "bigfig":
		fontLines = 4
	case "small":
		fontLines = 5	
	case "phoenix", "o2", "starwar", "stop":
		fontLines = 7
	case "standard", "shadow", "thinkertoy", "arob", "zigzag", "henry3D", "doom":
		fontLines = 8
	case "blocks":
		fontLines = 11
	}
}
