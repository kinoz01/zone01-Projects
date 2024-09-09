package asciiart

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

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

// IsValidColor converts any color format to an ANSI color code.
func IsValidColor(color string) string {
	if strings.HasPrefix(color, "rgb") {
		return ParseRGB(color)
	} else if strings.HasPrefix(color, "#") {
		return ParseHex(color)
	} else if strings.HasPrefix(color, "hsl") {
		return ParseHSL(color)
	} else {
		return ParseColorName(color)
	}
}

// case "black", "rgb(0, 0, 0)", "#000000", "hsl(0, 0%, 0%)":
// ParseRGB converts an RGB string to an ANSI color code.
func ParseRGB(color string) string {
	re := regexp.MustCompile(`\Argb\((\d+),\s*(\d+),\s*(\d+)\)$`)
	matches := re.FindStringSubmatch(color)
	if matches == nil {
		return ""
	}
	r, err1 := strconv.Atoi(matches[1])
	g, err2 := strconv.Atoi(matches[2])
	b, err3 := strconv.Atoi(matches[3])

	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 || err1 != nil || err2 != nil || err3 != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// ParseHex converts a Hex string to an ANSI color code.
func ParseHex(color string) string {
	color = strings.TrimPrefix(color, "#")

	if len(color) == 3 {
		color = string(color[0]) + string(color[0]) +
			string(color[1]) + string(color[1]) +
			string(color[2]) + string(color[2])
	} else if len(color) != 6 {
		return ""
	}
	r64, err1 := strconv.ParseInt(color[0:2], 16, 64)
	g64, err2 := strconv.ParseInt(color[2:4], 16, 64)
	b64, err3 := strconv.ParseInt(color[4:6], 16, 64)

	r, g, b := int(r64), int(g64), int(b64)
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 || err1 != nil || err2 != nil || err3 != nil {
		return ""
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// ParseHSL converts an HSL string to an ANSI color code.
func ParseHSL(color string) string {
	re := regexp.MustCompile(`\Ahsl\((\d+),\s*(\d+)%,\s*(\d+)%\)$`)
	matches := re.FindStringSubmatch(color)
	if matches == nil {
		return ""
	}
	h_i, err1 := strconv.Atoi(matches[1])
	s_i, err2 := strconv.Atoi(matches[2])
	l_i, err3 := strconv.Atoi(matches[3])

	if h_i < 0 || h_i > 360 || s_i < 0 || s_i > 100 || l_i < 0 || l_i > 100 || err1 != nil || err2 != nil || err3 != nil {
		return ""
	}

	h, s, l := float64(h_i)/360, float64(s_i)/100, float64(l_i)/100

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h*6, 2)-1))
	m := l - c/2

	var r64, g64, b64 float64
	switch {
	case h_i < 60:
		r64, g64, b64 = c, x, 0
	case h_i < 120:
		r64, g64, b64 = x, c, 0
	case h_i < 180:
		r64, g64, b64 = 0, c, x
	case h_i < 240:
		r64, g64, b64 = 0, x, c
	case h_i < 300:
		r64, g64, b64 = x, 0, c
	default:
		r64, g64, b64 = c, 0, x
	}
	r, g, b := int((r64+m)*255), int((g64+m)*255), int((b64+m)*255)

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// ParseColorName converts a color name to an ANSI color code.
func ParseColorName(color string) string {
	switch strings.ToLower(color) {
	case "black":
		return black
	case "red":
		return red
	case "green":
		return green
	case "yellow":
		return yellow
	case "blue":
		return blue
	case "magenta":
		return magenta
	case "cyan":
		return cyan
	case "white":
		return white
	case "sky":
		return sky
	case "orange":
		return orange
	case "forest":
		return forest
	case "ocean":
		return ocean
	case "lavender":
		return lavender
	case "rose":
		return rose
	case "lemon":
		return lemon
	case "turquoise":
		return turquoise
	case "cherry":
		return cherry
	case "emerald":
		return emerald
	default:
		return ""
	}
}
