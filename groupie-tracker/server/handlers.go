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

// Handle serving both CSS and JS content
func FileHandler(w http.ResponseWriter, r *http.Request) {
	var filePath string

	// Check if the request is for a CSS or JS file and serve accordingly
	if strings.HasPrefix(r.URL.Path, "/css/") {
		filePath = "frontend/css/" + strings.TrimPrefix(r.URL.Path, "/css/")
	} else if strings.HasPrefix(r.URL.Path, "/js/") {
		filePath = "frontend/js/" + strings.TrimPrefix(r.URL.Path, "/js/")
	}

	// Read the file from the filesystem
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		ErrorHandler(w, http.StatusForbidden, http.StatusText(http.StatusForbidden), "You don't have permission to access this link!")
		return
	}

	// Serve the file content
	http.ServeContent(w, r, filePath, time.Now(), bytes.NewReader(fileBytes))
}
