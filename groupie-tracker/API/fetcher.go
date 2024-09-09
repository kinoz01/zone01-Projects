package apiserver

import (
	"fmt"
	"groupie/server"
	"net/http"
	"sync"
)

func FetchData(id string, apiArtist *server.Artist, apiArtistLocations *server.Locations, apiArtistRelations *server.Relations, apidata *ApiData, w http.ResponseWriter) {
	var wg sync.WaitGroup
	var errCh = make(chan error, 4)

	wg.Add(4)

	// Fetch artist data concurrently
	go func() {
		defer wg.Done()
		err := server.FetchData(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id), apiArtist)
		if err != nil {
			errCh <- fmt.Errorf("failed to fetch artist: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := server.FetchData(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", id), apiArtistLocations)
		if err != nil {
			errCh <- fmt.Errorf("failed to fetch locations: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := server.FetchData(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id), apiArtistRelations)
		if err != nil {
			errCh <- fmt.Errorf("failed to fetch dates: %w", err)
		}
	}()

	// Initialise ApiData concurrently
	go func() {
		defer wg.Done()
		apidata.Initialise()
	}()

	wg.Wait()
	close(errCh)

	// Handle errors from any of the goroutines
	if len(errCh) > 0 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
