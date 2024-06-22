package api

import (
	"asciiArt/asciiart"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		NotFoundHandler(w)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Text   string
		Banner string
		Art    string
		Fonts  []string
	}{
		Text:   "Hello World!",
		Banner: "standard",
		Art:    "",
		Fonts: []string{
			"small", "phoenix", "o2", "starwar", "stop", "varsity", "standard",
			"shadow", "thinkertoy", "arob", "zigzag", "henry3D", "doom", "tiles",
			"jacky", "catwalk", "coins", "fire", "jazmine", "matrix", "blocks",
			"univers", "impossible", "georgi",
		},
	}

	tmpl.Execute(w, data)
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request: Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	art, err := asciiart.ASCIIArt(text, banner)
	if err != nil {
		http.Error(w, art, http.StatusInternalServerError)
		return
	}

	data := struct {
		Text   string
		Banner string
		Art    string
		Fonts  []string
	}{
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

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func NotFoundHandler(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	tmpl.Execute(w, nil)
}
