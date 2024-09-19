package apiserver

import (
	"groupie/server"
	"sync"
)

// Define a function that uses goroutines to fetch data concurrently
func SetData(apiArtist server.Artist, apiArtistRelations server.Relations, apiArtistLocations server.Locations, apidata ApiData) Data {
	var wg sync.WaitGroup
	var data Data

	// Channels to hold the results
	imageChan := make(chan string)
	membersImagesChan := make(chan map[string]string)
	locationsDatesChan := make(chan map[string][]string)
	youtubeLinksChan := make(chan []string)

	// Start the goroutines

	// Goroutine for Search
	wg.Add(1)
	go func() {
		defer wg.Done()
		imageChan <- Search(apiArtist.Name, "logo", apidata.BandImages)
	}()

	// Goroutine for GetMembersImages
	wg.Add(1)
	go func() {
		defer wg.Done()
		membersImagesChan <- GetMembersImages(apiArtist.Members, apidata.MembersImages)
	}()

	// Goroutine for GetLocationsDates
	wg.Add(1)
	go func() {
		defer wg.Done()
		locationsDatesChan <- GetLocationsDates(apiArtistRelations.DatesLocations, apiArtistLocations.Locations)
	}()

	// Goroutine for GetYoutubeLinks
	wg.Add(1)
	go func() {
		defer wg.Done()
		youtubeLinksChan <- GetYoutubeLinks(apiArtist.Name, apidata.YoutubeLinks)
	}()

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(imageChan)
		close(membersImagesChan)
		close(locationsDatesChan)
		close(youtubeLinksChan)
	}()

	// Collect results from the channels
	data.Image = <-imageChan
	data.MembersImages = <-membersImagesChan
	data.DatesLocations = <-locationsDatesChan
	data.YoutubeUrl = <-youtubeLinksChan

	return data
}
