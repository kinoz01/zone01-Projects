package server

import (
	"net/http"
)

// Check for quiting conditions in the Home request.
func CheckHomeRequest(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/" {
		ErrorHandler(w, 404, "Look like you're lost!", "The page you are looking for is not available!")
		return true
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, 405, http.StatusText(http.StatusMethodNotAllowed), "Only GET method is allowed!")
		return true
	}
	return false
}

// Check for quiting conditions in the Post request.
func CheckPostRequest(w http.ResponseWriter, r *http.Request) bool {

	// Limit the size of the request body
	r.Body = http.MaxBytesReader(w, r.Body, 30000)

	// Read form values after limiting the request body size
	if err := r.ParseForm(); err != nil {
		ErrorHandler(w, 413, "Request size exceeds the limit.", "Payload Too Large")
		return true
	}

	if r.Method != http.MethodPost {
		ErrorHandler(w, 405, http.StatusText(http.StatusMethodNotAllowed), "Only POST method is allowed!")
		return true
	}

	// Check form values and handle actions
	if CheckFormValues(w, r) {
		return true
	}

	return false
}


// Check the Form value and make sure to return an error when changed using inspect.
func CheckFormValues(w http.ResponseWriter, r *http.Request) bool {
	// Required form values
	requiredFields := []string{"banner", "text", "action"}
	for _, field := range requiredFields {
		if r.FormValue(field) == "" {
			ErrorHandler(w, 400, "Make sure your input is correct!", "Bad Request!")
			return true
		}
	}

	// Validate the action field
	action := r.FormValue("action")
	switch action {
	case "preview":
		PreviewFontsHandler(w, r, r.FormValue("text"))
		return true
	case "download":
		DownloadHandler(w, r, r.FormValue("text"), r.FormValue("banner"))
		return true
	case "generate":
		// continue
	default:
		ErrorHandler(w, 400, "Make sure your input is correct!", "Bad Request!")
		return true
	}

	return false
}
