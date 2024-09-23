package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type ErrorData struct {
	Msg1       string
	Msg2       string
	StatusCode int
}

// Parse and execute error html page depending on error type.
func ErrorHandler(w http.ResponseWriter, statusCode int, msg1, msg2 string, err error) {

	// print errors in case of intenal server error
	if err != nil && statusCode == 500 {
		log.Println(err)
	}

	Error := ErrorData{
		Msg1:       msg1,
		Msg2:       msg2,
		StatusCode: statusCode,
	}

	tmpl, err := template.ParseFiles("frontend/templates/error.html")
	if err != nil {
		ServeRentryError(w, err, msg1, msg2, statusCode)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, Error); err != nil {
		ServeRentryError(w, err, msg1, msg2, statusCode)
		return
	}

	w.WriteHeader(statusCode)
	// If successful, write the buffer content to the ResponseWriter
	buf.WriteTo(w)
}

func ServeRentryError(w http.ResponseWriter, err error, msg1, msg2 string, statusCode int) {
	log.Println(err)
	errBody, err := GetErrorPage()
	if err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}
	// Replace placeholders in the error page with dynamic messages
	errBody = strings.ReplaceAll(errBody, "{{.Msg1}}", msg1)
	errBody = strings.ReplaceAll(errBody, "{{.Msg2}}", msg2)
	errBody = strings.ReplaceAll(errBody, "{{.StatusCode}}", strconv.Itoa(statusCode))
	if statusCode == 500 {
		errBody = strings.ReplaceAll(errBody, `<a href="/"><button class="submit">Go to Home</button>`, "")
	}

	// Set the response header and write the error page
	w.WriteHeader(statusCode)
	w.Write([]byte(errBody))
}

// GetErrorPage retrieves the error page from rentry.co based on the error code.
func GetErrorPage() (string, error) {
	// Build the URL based on the error number
	url := APILinks.ErrorPage

	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch error page: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is not OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch error page: received status %d", resp.StatusCode)
	}

	// Read the HTML content from the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read error page content: %v", err)
	}

	// Convert the body to a string and return
	return string(body), nil
}
