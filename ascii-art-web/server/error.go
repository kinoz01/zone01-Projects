package server

import (
	"bytes"
	"net/http"
	"text/template"
)

type ErrorData struct {
	Msg1        string
	Msg2        string
	StatusCode  int
	HomeAddress string
}

// Parse and execute error html page depending on error type.
func ErrorHandler(w http.ResponseWriter, statusCode int, msg1, msg2 string) {

	Error := ErrorData{
		Msg1:       msg1,
		Msg2:       msg2,
		StatusCode: statusCode,
	}
	Error.HomeAddress = "http://127.0.0.1:" + Port

	w.WriteHeader(Error.StatusCode)

	tmpl, err := template.ParseFS(TemplatesFS, "frontend/templates/error.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, Error); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// If successful, write the buffer content to the ResponseWriter
	_, err = buf.WriteTo(w)
	if err != nil {
		// Handle any error that might occur while writing the buffer to the response
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
