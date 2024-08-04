package api

import (
	"fmt"
	"net/http"
	"text/template"
)

type ErrorData struct {
	Msg1       string
	Msg2       string
	StatusCode int
}

// Page not found.
func Error404(w http.ResponseWriter) {
	Error := ErrorData{
		Msg1:       "Look like you're lost!",
		Msg2:       "The page you are looking for is not available!",
		StatusCode: 404,
	}
	ErrorHandler(w, Error)
}

// Method not allowed.
func Error405(w http.ResponseWriter, method string) {
	msg2 := fmt.Sprintf("Only %s method is allowed!", method)
	Error := ErrorData{
		Msg1:       http.StatusText(http.StatusMethodNotAllowed),
		Msg2:       msg2,
		StatusCode: http.StatusMethodNotAllowed,
	}
	ErrorHandler(w, Error)
}

// Internal server error.
func Error500(w http.ResponseWriter) {
	Error := ErrorData{
		Msg1:       "Something seems wrong, try again later!",
		Msg2:       "Internal Server Error!",
		StatusCode: http.StatusInternalServerError,
	}
	ErrorHandler(w, Error)
}

// Bad request.
func Error400(w http.ResponseWriter) {
	Error := ErrorData{
		Msg1:       "Make sure your input is correct!",
		Msg2:       "Bad Request!",
		StatusCode: http.StatusBadRequest,
	}
	ErrorHandler(w, Error)
}

// Parse and execute error html page depending on error type.
func ErrorHandler(w http.ResponseWriter, ErrData ErrorData) {

	w.WriteHeader(ErrData.StatusCode)

	tmpl, err := template.ParseFS(TemplatesFs, "templates/error.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, ErrData); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
