package apiserver

import (
	"encoding/json"
	"fmt"
	"groupie/server"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ImageResult struct {
	Position       int     `json:"position"`
	Original       string  `json:"original"`
	OriginalWidth  float64 `json:"original_width"`
	OriginalHeight float64 `json:"original_height"`
}

type SearchResults struct {
	Images []ImageResult `json:"images_results"`
}

// Cache structure to store query and image URL pairs
type Cache map[string]string

const (
	apiKey    = "00f93dbcc368abe0041df6b8f4ff6b90ab59665e2f5867cf6ff0ee74f9e15db9" // serpapi key
	cacheFile = "./API/locations.json" // json file name
)

func GetApiImage(query string) string {
	
	query = FormatQuery(query)

	// Load cache from file or initialize a new one
	cache, err := LoadCache()
	if err != nil {
		cache = make(Cache)
	}

	// Check if the query already exists in the cache
	if url, found := cache[query]; found {
		return url
	}

	apiUrl := fmt.Sprintf(server.APILinks.SerpApi, query, apiKey)

	// Send HTTP GET request
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Println("failed to fetch the JSON: ", err)
		return ""
	}
	defer resp.Body.Close()

	// Decode the JSON data
	var searchResults SearchResults
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&searchResults); err != nil {
		log.Println("failed to decode JSON: ", err)
		return ""
	}

	// Process the fetched results and find the image URL
	imageURL, err := ProcessImageResults(searchResults)
	if err != nil {
		return ""
	}

	// Save the result to the cache
	cache[query] = imageURL
	if err := SaveCache(cache); err != nil {
		log.Printf("Failed to save cache: %v", err)
	}

	return imageURL
}

// processImageResults processes the search results and returns the appropriate image URL
func ProcessImageResults(searchResults SearchResults) (string, error) {
	// Iterate over positions 0 to 10 to find a matching image
	for pos := 0; pos <= 20; pos++ {
		for _, result := range searchResults.Images {
			if result.Position == pos && result.OriginalWidth >= 1200 && result.OriginalWidth >= 1.5*result.OriginalHeight {
				return result.Original, nil
			}
		}
	}

	// If no suitable image is found, return an error
	return "", fmt.Errorf("no valid image found")
}

func FormatQuery(query string) string {
	result := ""
	for _, char := range query {
		if char == ' ' {
			result += "+"
		} else if char == ',' {
			result += "%2C"
		} else {
			result += string(char)
		}
	}
	return result
}

// loadCache loads the cache from the cache file
func LoadCache() (Cache, error) {
	if _, err := os.Stat(cacheFile); err == nil {
		// Cache file exists, read it
		data, err := ioutil.ReadFile(cacheFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read cache file: %v", err)
		}

		// Parse the cache data
		var cache Cache
		if err := json.Unmarshal(data, &cache); err != nil {
			return nil, fmt.Errorf("failed to parse cache data: %v", err)
		}

		return cache, nil
	}
	// Cache file does not exist, return an empty cache
	return make(Cache), nil
}

// saveCache saves the cache to the cache file
func SaveCache(cache Cache) error {
	cacheData, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache data: %v", err)
	}
	if err := ioutil.WriteFile(cacheFile, cacheData, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %v", err)
	}
	return nil
}
