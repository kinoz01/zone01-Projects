package wget

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

// Prepares the HTTP client with custom redirect handling.
func PrepareHTTPClient() *http.Client {
	return &http.Client{
		Timeout:       0,
		CheckRedirect: RedirectHandler,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			ForceAttemptHTTP2: false,
		},
	}
}

// RedirectHandler handles redirects and logs them
func RedirectHandler(req *http.Request, via []*http.Request) error {
	if len(via) >= 10 {
		return fmt.Errorf("maximum redirects reached")
	}

	// Get the response that caused the redirect
	if resp := req.Response; resp != nil {
		fmt.Fprintf(LogOutput, "\nHTTP request sent, awaiting response... %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		location := req.URL.String()
		fmt.Fprintf(LogOutput, "Location: %s [following]\n", location)
	}
	return nil
}
