package server

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	err := FetchData(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id), &artist)
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
			(*artists)[i].Image = "https://i.postimg.cc/hjXqxwCS/500x500.jpg"
		case "Thirty Seconds to Mars":
			(*artists)[i].Image = "https://i.postimg.cc/J7jQbWcT/pngegg.png"
		case "Eminem":
			(*artists)[i].Image = "https://i.postimg.cc/gkg20Qyf/eminem.png"
		}
	}
}

// Initialise ports (api & application ports)
func (p *Ports) InitialisePorts() {
	p.Port = os.Getenv("PORT")
	if p.Port == "" {
		p.Port = "8088" // Default port
	}

	p.ApiPort = os.Getenv("ApiPORT")
	if p.ApiPort == "" {
		p.ApiPort = "4000" // Default port
	}
}
