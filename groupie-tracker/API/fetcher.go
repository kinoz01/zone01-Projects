package apiserver

import (
	"fmt"
	"groupie/server"
	"net/http"
	"sync"
)

// Use Goroutines to fetch herokuapp artists data and decode it in input structs, while also intialising an input ApiDAta struct in the data.go page.
func FetchData(id string, apiArtist *server.Artist, apiArtistLocations *server.Locations, apiArtistRelations *server.Relations, apidata *ApiData, w http.ResponseWriter) {

	var wg sync.WaitGroup
	var errCh = make(chan error, 4)
	wg.Add(4)

	// Fetch artist data concurrently
	go func() {
		defer wg.Done()
		err := server.FetchData(fmt.Sprintf(server.APILinks.Artist, id), apiArtist)
		if err != nil {
			errCh <- fmt.Errorf("failed to fetch artist: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := server.FetchData(fmt.Sprintf(server.APILinks.Locations, id), apiArtistLocations)
		if err != nil {
			errCh <- fmt.Errorf("failed to fetch locations: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := server.FetchData(fmt.Sprintf(server.APILinks.Relations, id), apiArtistRelations)
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
