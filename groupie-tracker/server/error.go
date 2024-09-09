package server

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"text/template"
)

type ErrorData struct {
	Msg1        string
	Msg2        string
	StatusCode  int
}

// Parse and execute error html page depending on error type.
func ErrorHandler(w http.ResponseWriter, statusCode int, msg1, msg2 string) {

	Error := ErrorData{
		Msg1:       msg1,
		Msg2:       msg2,
		StatusCode: statusCode,
	}

	w.WriteHeader(Error.StatusCode)

	tmpl, err := template.ParseFiles("frontend/templates/error.html")
	if err != nil {
		PrintLog(err)
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, Error); err != nil {
		PrintLog(err)
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}
	// If successful, write the buffer content to the ResponseWriter
	buf.WriteTo(w)
}

// Print logs depending on environment variable.
func PrintLog(err error) {
	Logs = os.Getenv("LOG")
	if Logs == "set" {
		log.Println(err)
		return
	}
}
