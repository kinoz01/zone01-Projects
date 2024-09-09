package server

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := CheckHomeRequest(w, r); err {
		return
	}

	var artists []Artist
	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		PrintLog(err)
		ErrorHandler(w, http.StatusInternalServerError, "Failed to fetch artists", "Internal Server Error!")
		return
	}
	ReplaceImages(&artists)

	ParseAndExecute(w, artists, "frontend/templates/index.html")
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if err := CheckArtistRequest(w, r, id); err {
		return
	}

	var artistDetails ArtistDetails
	if quit := FetchSyncData(w, id, &artistDetails); quit {
		return
	}
	ParseAndExecute(w, artistDetails, "frontend/templates/artist.html")
}

// Handle serving css content, while blocking access to paths "/css/..."
func CSSHandler(w http.ResponseWriter, r *http.Request) {
	// Strip the "/css/" prefix from the URL path to get the relative file path
	filePath := "frontend/css/" + strings.TrimPrefix(r.URL.Path, "/css/")
	// Read the file from the embedded filesystem
	cssBytes, err := os.ReadFile(filePath)
	if err != nil {
		ErrorHandler(w, http.StatusForbidden, http.StatusText(http.StatusForbidden), "You don't have permission to access this link!")
		return
	}
	// Serve the file content
	http.ServeContent(w, r, filePath, time.Now(), bytes.NewReader(cssBytes))
}

// Handle serving JavaScript content, while blocking access to paths "/js/..."
func JSHandler(w http.ResponseWriter, r *http.Request) {
	// Strip the "/js/" prefix from the URL path to get the relative file path
	filePath := "frontend/js/" + strings.TrimPrefix(r.URL.Path, "/js/")
	// Read the file from the embedded filesystem
	jsBytes, err := os.ReadFile(filePath)
	if err != nil {
		ErrorHandler(w, http.StatusForbidden, http.StatusText(http.StatusForbidden), "You don't have permission to access this link!")
		return
	}
	// Serve the file content
	http.ServeContent(w, r, filePath, time.Now(), bytes.NewReader(jsBytes))
}
