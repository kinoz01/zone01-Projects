package api

import (
	"asciiArt/asciiart"
	"embed"
	"net/http"
	"text/template"
)

var TemplateFs embed.FS

type WebPageData struct {
	Text   string
	Banner string
	Art    string
	Fonts  []string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		Error404(w)
		return
	}

	if r.Method != http.MethodGet {
		Error405(w, "GET")
		return
	}

	tmpl, err := template.ParseFS(TemplateFs, "templates/index.html")
	if err != nil {
		Error500(w)
		return
	}

	data := WebPageData{
		Text:   "Type Something!",
		Banner: "standard",
		Art:    "",
		Fonts: []string{
			"small", "phoenix", "o2", "starwar", "stop", "varsity", "standard",
			"shadow", "thinkertoy", "arob", "zigzag", "henry3D", "doom", "tiles",
			"jacky", "catwalk", "coins", "fire", "jazmine", "matrix", "blocks",
			"univers", "impossible", "georgi",
		},
	}

	if err := tmpl.Execute(w, data); err != nil {
		Error500(w)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Error405(w, "POST")
		return
	}

	if err := r.ParseForm(); err != nil {
		Error400(w)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner == "" {
		Error400(w)
		return
	}

	art, err := asciiart.ASCIIArt(text, banner)
	if err != nil {
		Error400(w)
		return
	}

	data := WebPageData{
		Text:   text,
		Banner: banner,
		Art:    art,
		Fonts: []string{
			"small", "phoenix", "o2", "starwar", "stop", "varsity", "standard",
			"shadow", "thinkertoy", "arob", "zigzag", "henry3D", "doom", "tiles",
			"jacky", "catwalk", "coins", "fire", "jazmine", "matrix", "blocks",
			"univers", "impossible", "georgi",
		},
	}

	tmpl, err := template.ParseFS(TemplateFs, "templates/index.html")
	if err != nil {
		Error500(w)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		Error500(w)
		return
	}
}
