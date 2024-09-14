package server

import (
	"bytes"
	"log"
	"net/http"
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

	w.WriteHeader(Error.StatusCode)

	tmpl, err := template.ParseFiles("frontend/templates/error.html")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, Error); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}
	// If successful, write the buffer content to the ResponseWriter
	buf.WriteTo(w)
}

