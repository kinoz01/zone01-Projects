package asciiart

// Here We define colors supported by the program.
const (
	reset     = "\033[0m"
	black     = "\033[30m"
	red       = "\033[31m"
	green     = "\033[32m"
	yellow    = "\033[33m"
	blue      = "\033[34m"
	magenta   = "\033[35m"
	cyan      = "\033[36m"
	white     = "\033[37m"
	sky       = "\033[38;5;111m"
	orange    = "\033[38;5;208m"
	forest    = "\033[38;5;28m"
	ocean     = "\033[38;5;27m"
	lavender  = "\033[38;5;183m"
	rose      = "\033[38;5;197m"
	lemon     = "\033[38;5;226m"
	turquoise = "\033[38;5;80m"
	cherry    = "\033[38;5;161m"
	emerald   = "\033[38;5;46m"
)

// IsValidColor will link a string color with an ansi color code. In case of invalid color it will return an empty string.
func IsValidColor(color string) (ansiColor string) {
	switch color {
	case "black", "rgb(0, 0, 0)", "#000000", "hsl(0, 0%, 0%)":
		ansiColor = black
	case "red", "rgb(255, 0, 0)", "#ff0000", "hsl(0, 100%, 50%)":
		ansiColor = red
	case "green", "rgb(0, 255, 0)", "#00ff00", "hsl(120, 100%, 50%)":
		ansiColor = green
	case "yellow", "rgb(255, 255, 0)", "#ffff00", "hsl(60, 100%, 50%)":
		ansiColor = yellow
	case "blue", "rgb(0, 0, 255)", "#0000ff", "hsl(240, 100%, 50%)":
		ansiColor = blue
	case "magenta", "rgb(255, 0, 255)", "#ff00ff", "hsl(300, 100%, 50%)":
		ansiColor = magenta
	case "cyan", "rgb(0, 255, 255)", "#00ffff", "hsl(180, 100%, 50%)":
		ansiColor = cyan
	case "white", "rgb(255, 255, 255)", "#ffffff", "hsl(0, 0%, 100%)":
		ansiColor = white
	case "sky":
		ansiColor = sky
	case "orange":
		ansiColor = orange
	case "forest":
		ansiColor = forest
	case "ocean":
		ansiColor = ocean
	case "lavender":
		ansiColor = lavender
	case "rose":
		ansiColor = rose
	case "lemon":
		ansiColor = lemon
	case "turquoise":
		ansiColor = turquoise
	case "cherry":
		ansiColor = cherry
	case "emerald":
		ansiColor = emerald
	}
	return ansiColor
}
