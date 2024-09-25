package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"text/template"
)

// Check for quiting conditions in the Home request.
func CheckHomeRequest(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/" {
		ErrorHandler(w, 404, "Look like you're lost!", "The page you are looking for is not available!", nil)
		return true
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, 405, http.StatusText(http.StatusMethodNotAllowed), "Only GET method is allowed!", nil)
		return true
	}
	return false
}

// Check for quiting conditions in the artist path request.
func CheckArtistRequest(w http.ResponseWriter, r *http.Request, id string) bool {
	if r.Method != http.MethodGet {
		ErrorHandler(w, 405, http.StatusText(http.StatusMethodNotAllowed), "Only GET method is allowed!", nil)
		return true
	}
	if len(r.URL.Query()) == 0 || len(r.URL.Query()) > 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return true
	}
	if _, err := strconv.Atoi(id); err != nil {
		ErrorHandler(w, 404, "Look like you're lost!", "The page you are looking for is not available!", err)
		return true
	}

	var artist Artist
	err := FetchData(fmt.Sprintf(APILinks.Artist, id), &artist)
	if err != nil || artist.ID == 0 {
		ErrorHandler(w, 404, "Look like you're lost!", "The page you are looking for is not available!", err)
		return true
	}
	return false
}

// Parse the html files and excute them after checking for errors.
func ParseAndExecute(w http.ResponseWriter, data any, file string) {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Something seems wrong, try again later!", "Internal Server Error!", err)
		return
	}

	// write to a temporary buffer instead of writing directly to w.
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		ErrorHandler(w, 500, "Something seems wrong, try again later!", "Internal Server Error!", err)
		return
	}
	// If successful, write the buffer content to the ResponseWriter
	buf.WriteTo(w)
}

// Replace some default api images with other choosen images.
func ReplaceImages(artists *[]Artist) {
	for i, artist := range *artists {
		switch artist.Name {
		case "Mamonas Assassinas":
			(*artists)[i].Image = APILinks.ReplacedImages[0]
		case "Thirty Seconds to Mars":
			(*artists)[i].Image = APILinks.ReplacedImages[1]
		case "Eminem":
			(*artists)[i].Image = APILinks.ReplacedImages[2]
		}
	}
}

// Initialise ports (api & application ports) using environement variables.
// For example use export PORT=:$PORT to set port where the server should start.
func (p *Ports) InitialisePorts() {

	p.Port = os.Getenv("PORT")
	p.ApiPort = os.Getenv("APIPORT")

}

// Initialise the APILinks global struct using apiLinks.json.
func InitialiseApiLinks() {
	// Load the JSON with Heroku API links and other links.
	var err error
	APILinks, err = LoadApiLinks("./server/apiLinks.json")
	if err != nil {
		log.Fatalf("Error loading API links: %v", err)
	}
}

func LoadApiLinks(filename string) (*ApiLinks, error) {
	// Load the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into ApiLinks struct
	var links ApiLinks
	err = json.Unmarshal(content, &links)
	if err != nil {
		return nil, err
	}

	return &links, nil
}

func GetAllPlacesNames(artists *[]Artist) {
	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	for i := range *artists {
		artist := &(*artists)[i]

		// Start a new goroutine for each artist
		wg.Add(1)
		go func(artist *Artist) {
			defer wg.Done() // Mark this goroutine as done when it finishes

			var location Locations

			// Fetch data for this artist's locations
			err := FetchData(artist.Locations, &location)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Append the fetched locations to the artist's LOCATIONS field
			artist.LOCATIONS = append(artist.LOCATIONS, location.Locations...)
		}(artist)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
