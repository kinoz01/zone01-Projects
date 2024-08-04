package api

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"asciiArt/asciiart"
)

var TemplatesFs embed.FS

type WebPageData struct {
	Text   string
	Banner string
	Art    string
	Fonts  []string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error404(w)
		return
	}

	if r.Method != http.MethodGet {
		Error405(w, "GET")
		return
	}

	tmpl, err := template.ParseFS(TemplatesFs, "templates/index.html")
	if err != nil {
		Error500(w)
		return
	}

	data := WebPageData{
		Text:   "Hello World!",
		Banner: "standard",
		Art:    "",
	}
	data.ReadFonts()
	data.ReadUserFonts()
	if data.Fonts == nil {
		Error500(w)
		return
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

	action := r.FormValue("action")
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if action == "preview" {
		PreviewFontsHandler(w, r)
		return
	}

	if banner == "" {
		Error400(w)
		return
	}

	art, err := asciiart.ASCIIArt(text, banner)
	if err != nil {
		Error400(w)
		return
	}

	if action == "download" {
		DownloadHandler(w, r, art)
		return
	}

	data := WebPageData{
		Text:   text,
		Banner: banner,
		Art:    art,
	}
	data.ReadFonts()
	data.ReadUserFonts()

	tmpl, err := template.ParseFS(TemplatesFs, "templates/index.html")
	if err != nil {
		Error500(w)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		Error500(w)
		return
	}
}

func (d *WebPageData) ReadFonts() {
	BannersFS := asciiart.Banners
	entries, err := BannersFS.ReadDir("banners")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".txt") {
			d.Fonts = append(d.Fonts, strings.TrimSuffix(entry.Name(), ".txt"))
		}
	}
}

func (d *WebPageData) ReadUserFonts() {
	files, err := os.ReadDir("banners")
	if err != nil {
		return
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			d.Fonts = append(d.Fonts, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}
}

func PreviewFontsHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFS(TemplatesFs, "templates/preview.html")
	if err != nil {
		Error500(w)
		return
	}
	art, err := asciiart.ASCIIArt("Hello World!", "shadow")
	if err != nil {
		Error400(w)
		return
	}

	if err := templ.Execute(w, art); err != nil {
		Error500(w)
		return
	}
}

func DownloadHandler(w http.ResponseWriter, r *http.Request, art string) {

	// Set content type and disposition headers
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii.txt")

	// Write the string to the response
	_, err := w.Write([]byte(art))
	if err != nil {
		Error500(w)
		return
	}
}
