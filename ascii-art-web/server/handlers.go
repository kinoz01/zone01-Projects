package server

import (
	"bytes"
	"embed"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"asciiArt/asciiart"
)

var TemplatesFS embed.FS
var CssFS embed.FS
var Style string

type WebPageData struct {
	Text   string
	Banner string
	Art    string
	Fonts  []string
}

// Initialize the fields of an instance of the WebPageData struct into needed values.
func NewWebPagedata(text, banner, art string) WebPageData {
	data := WebPageData{
		Text:   text,
		Banner: banner,
		Art:    art,
	}
	data.ReadFonts("myFonts")
	data.ReadFonts("userFonts")

	return data
}

// Handle home path (parse and execute the index.html)
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var tmpl *template.Template
	var err error

	if err := CheckHomeRequest(w, r); err {
		return
	}

	if Style == "K"  || StyleEnv == "K" {
		tmpl, err = template.ParseFS(TemplatesFS, "frontend/templates/index2.html")
	} else {
		tmpl, err = template.ParseFS(TemplatesFS, "frontend/templates/index.html")
	}

	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Something seems wrong, try again later!", "Internal Server Error!")
		return
	}

	data := NewWebPagedata("Hello World!", "standard", "")
	if data.Fonts == nil {
		ErrorHandler(w, http.StatusInternalServerError, "Something seems wrong, try again later!", "Internal Server Error!")
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		ErrorHandler(w, 500, "Something seems wrong, try again later!", "Internal Server Error!")
		return
	}
	// If successful, write the buffer content to the ResponseWriter
	_, err = buf.WriteTo(w)
	if err != nil {
		// Handle any error that might occur while writing the buffer to the response
		ErrorHandler(w, 500, "Something seems wrong, try again later!", "Internal Server Error!")
	}
}

// Handle ascii-art path (parse and execute the index.html but with new WebPageData after receiving the Post data from the user)
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {

	var tmpl *template.Template
	var err error

	if quit := CheckPostRequest(w, r); quit {
		return
	}

	text := "\r\n" + r.FormValue("text")
	banner := r.FormValue("banner")

	art, err := asciiart.ASCIIArt(text, banner)
	if err != nil {
		ErrorHandler(w, 400, "Make sure your input is correct!", "Bad Request!")
		return
	}

	data := NewWebPagedata(text, banner, art)

	if Style == "K" || StyleEnv == "K" {
		tmpl, err = template.ParseFS(TemplatesFS, "frontend/templates/index2.html")
	} else {
		tmpl, err = template.ParseFS(TemplatesFS, "frontend/templates/index.html")
	}

	if err != nil {
		ErrorHandler(w, 500, "Something seems wrong, try again later!", "Internal Server Error!")
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, 500, "Something seems wrong, try again later!", "Internal Server Error!")
		return
	}
}

// Handle Preview action button by parsing and executing preview.html with a map of fonts names and its corresponding art.
func PreviewFontsHandler(w http.ResponseWriter, r *http.Request, text string) {
	templ, err := template.ParseFS(TemplatesFS, "frontend/templates/preview.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Something seems wrong, try again later!", "Internal Server Error!")
		return
	}

	data := NewWebPagedata("", "", "")
	fontTextMap := make(map[string]string)

	for _, font := range data.Fonts {
		art, err := asciiart.ASCIIArt(text, font)
		if err != nil || art == "Non-ASCII characters aren't supported.\n" {
			ErrorHandler(w, http.StatusBadRequest, "Make sure your input is correct!", "Bad Request!")
			return
		}
		fontTextMap[font] = art
	}

	if err := templ.Execute(w, fontTextMap); err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Something seems wrong, try again later!", "Internal Server Error!")
		return
	}
}

// Handle Download action button in the Post to ascci-art.
func DownloadHandler(w http.ResponseWriter, r *http.Request, text, banner string) {
	art, erro := asciiart.ASCIIArt(text, banner)
	if erro != nil {
		ErrorHandler(w, 400, "Make sure your input is correct!", "Bad Request!")
		return
	}

	// Set content type, length, and disposition headers
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(art)))
	w.Header().Set("Content-Disposition", "attachement; filename=\"ascii.txt\"")

	// Write the string to the response
	_, err := w.Write([]byte(art))
	if err != nil {
		ErrorHandler(w, 500, "Something seems wrong, try again later!", "Internal Server Error!")
		return
	}
}

// Append fonts found in the banner embeded directory to Fonts field of WebPageData struct.
// And in case of userFonts append fonts found in an outside banner directory to Fonts field of WebPageData struct.
func (d *WebPageData) ReadFonts(directory string) {

	var files []fs.DirEntry
	var err error

	if directory == "myFonts" {
		files, err = asciiart.Banners.ReadDir("banners")
	} else if directory == "userFonts" {
		files, err = os.ReadDir("banners")
	}

	if err != nil {
		return
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			d.Fonts = append(d.Fonts, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}
}

// Handle serving css content, while blocking access to paths "/css/..."
func CSSHandler(w http.ResponseWriter, r *http.Request) {

	// Strip the "/css/" prefix from the URL path to get the relative file path
	filePath := "frontend/css/" + strings.TrimPrefix(r.URL.Path, "/css/")
	// Read the file from the embedded filesystem
	cssBytes, err := CssFS.ReadFile(filePath)
	if err != nil {
		ErrorHandler(w, http.StatusForbidden, http.StatusText(http.StatusForbidden), "You don't have permission to access this link!")
		return
	}

	// Serve the file content
	http.ServeContent(w, r, filePath, time.Now(), bytes.NewReader(cssBytes))
}
