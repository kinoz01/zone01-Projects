package main

import (
	"asciiArt/server"
	"embed"
)

//go:embed frontend
var templateFs embed.FS

//go:embed frontend/css
var cssFiles embed.FS

// go build -ldflags "-X 'main.Style=K'" . (choose between A and k style)
var Style = "A"

func main() {
    server.TemplatesFS = templateFs
	server.Style = Style
	server.CssFS = cssFiles
	server.NewServer()
}
