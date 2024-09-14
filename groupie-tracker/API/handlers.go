package apiserver

import (
	"encoding/json"
	"groupie/server"
	"net/http"
)

type Data struct {
	Image          string              `json:"image"`
	MembersImages  map[string]string   `json:"membersImages"`
	DatesLocations map[string][]string `json:"datesLocations"`
	YoutubeUrl     []string            `json:"youtubeUrl"`
}

// handle the /groupie path to response with an encoded json.
func DataHandler(w http.ResponseWriter, r *http.Request) {
	var apiArtist server.Artist
	var apiArtistLocations server.Locations
	var apiArtistRelations server.Relations
	var apidata ApiData

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	FetchData(id, &apiArtist, &apiArtistLocations, &apiArtistRelations, &apidata, w)
	data := SetData(apiArtist, apiArtistRelations, apiArtistLocations, apidata)

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the struct into JSON and serve it
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
