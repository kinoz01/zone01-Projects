package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Use FetchData function simultaneously using go routines, and retrieving the returned errors using channels.
func FetchSyncData(w http.ResponseWriter, id string, artistDetails *ArtistDetails) bool {

	var wg sync.WaitGroup
	t := make(chan error, 5)

	FetchDataChan := func(url string, data any) {
		defer wg.Done()
		err := FetchData(url, data)
		if err != nil {
			t <- err
			return
		}
	}
	wg.Add(5)
	//fmt.Println()
	go FetchDataChan(fmt.Sprintf(APILinks.Artist, id), &artistDetails.Artist)
	go FetchDataChan(fmt.Sprintf(APILinks.Locations, id), &artistDetails.Locations)
	go FetchDataChan(fmt.Sprintf(APILinks.Dates, id), &artistDetails.Dates)
	go FetchDataChan(fmt.Sprintf(APILinks.Relations, id), &artistDetails.Relations)
	go FetchDataChan(fmt.Sprintf("http://127.0.0.1:%s/groupie?id=%s", Port.ApiPort, id), &artistDetails.MyAPI)
	wg.Wait()
	close(t)
	if len(t) != 0 {
		ErrorHandler(w, http.StatusInternalServerError, "Failed to fetch api data", "Internal Server Error!", nil)
		return true
	}
	return false
}

// Decode json data from url into the struct interface.
func FetchData(url string, struc interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching data: ", err)
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(struc); err != nil {
		log.Println("Error decoding data: ", err)
		return err
	}
	return nil
}
