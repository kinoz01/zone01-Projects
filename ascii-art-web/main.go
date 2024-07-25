package main

import (
	"asciiArt/api"
	"embed"
)

//go:embed templates/*
var templateFs embed.FS

func main() {
    api.TemplateFs = templateFs
	api.NewServer()
}
