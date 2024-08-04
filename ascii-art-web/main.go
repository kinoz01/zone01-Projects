package main

import (
	"asciiArt/api"
	"embed"
)

//go:embed templates/*
var templateFs embed.FS

func main() {
    api.TemplatesFs = templateFs
	api.NewServer()
}
