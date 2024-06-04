package main

import (
	"asciiArt/asciiart"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
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
	}{
		Text:   "",
		Banner: "standard",
		Art:    "",
	}

	tmpl.Execute(w, data)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Text   string
		Banner string
		Art    string
	}{
		Text:   text,
		Banner: banner,
		Art:    art,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
