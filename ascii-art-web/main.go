package main

import (
	"html/template"
	"log"
	"net/http"
	"asciiArt/asciiart"
)


func main() {
    http.HandleFunc("/", handler)
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
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

    if r.Method == http.MethodPost {
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

        data = struct {
            Text   string
            Banner string
            Art    string
        }{
            Text:   text,
            Banner: banner,
            Art:    art,
        }
    }

    tmpl.Execute(w, data)
}