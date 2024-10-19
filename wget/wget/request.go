package wget

import (
	"fmt"
	"net/http"
	"net/url"
)

// Prepares the HTTP request with appropriate headers.
func PrepareHTTPRequest(normalizedURL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", normalizedURL, nil)
	if err != nil {
		fmt.Fprintf(LogOutput, "failed to create request: %v\n", err)
		return nil, err
	}

	// Set headers to mimic wget
	req.Header.Set("User-Agent", "Wget/1.21.1")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Connection", "Keep-Alive")

	return req, nil
}

// Executes the HTTP request and handles redirects.
func PerformRequest(req *http.Request, client *http.Client, normalizedURL string) (*http.Response, error) {
	var resp *http.Response
	var err error

	parsedURL, _ := url.Parse(normalizedURL)
	hostname := parsedURL.Hostname()

	fmt.Fprintf(LogOutput, "Connecting to %s (%s)... ", hostname, parsedURL.Host)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Fprintf(LogOutput, "failed: %v\n", err)
		return nil, err
	}
	fmt.Fprintf(LogOutput, "connected.\n")

	fmt.Fprintf(LogOutput, "HTTP request sent, awaiting response... %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))

	return resp, nil
}
