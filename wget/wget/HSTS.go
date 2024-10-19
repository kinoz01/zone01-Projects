package wget

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Add http:// or https:// to the url if missing.
// Print message if response forced to use https://
func HSTSurlCheck(rawURL string) (string, error) {
	normalizedURL, HSTS, err := NormalizeURL(rawURL)
	if err != nil {
		errorUrl := strings.TrimPrefix(rawURL, "http://")
		if strings.Contains(err.Error(), "no such host") {
			fmt.Fprintf(LogOutput, "Resolving %s (%s)... failed: Name or service not known.\nwget: unable to resolve host address ‘%s’\n", errorUrl, errorUrl, errorUrl)
		} else {
			fmt.Fprintf(LogOutput, "Invalid URL '%s': %v\n", errorUrl, err)
		}
		return "", err
	}

	if HSTS {
		fmt.Fprintf(LogOutput, "URL transformed to HTTPS due to an HSTS policy\n")
	}
	return normalizedURL, nil
}

// Add http:// to the url Parse it and Send a Head Request to the url to 
// check if it implement HSTS policy, if so return true and add https:// instead.
func NormalizeURL(rawURL string) (string, bool, error) {
	HSTS := false
	rawURL = strings.TrimSpace(rawURL)

	if strings.HasPrefix(rawURL, "https://") {
		return rawURL, false, nil
	}

	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "http://" + rawURL
	}

	// Parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", false, fmt.Errorf("invalid URL: %v", err)
	}

	// Use the default HTTP client to send a HEAD request
	resp, err := http.Head(parsedURL.String())
	if err != nil {
		return "", false, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	// If the server enforces HSTS, switch to HTTPS
	if hsts := resp.Header.Get("Strict-Transport-Security"); hsts != "" {
		parsedURL.Scheme = "https"
		HSTS = true
	}

	return parsedURL.String(), HSTS, nil
}
