package asciiart

var fontLines int

func InitFontLines(font string) {
	switch font {
	case "small":
		fontLines = 5	
	case "phoenix":
		fontLines = 7
	case "standard", "shadow", "thinkertoy", "arob", "zigzag":
		fontLines = 8
	case "blocks":
		fontLines = 11
	}
}
