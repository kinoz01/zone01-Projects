package wget

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Handle client, request, response and content saving (basically everything)
func Wget(rawURL string) error {

	normalizedURL, err := HSTSurlCheck(rawURL)
	if err != nil {
		return err
	}

	timestamp := time.Now().Format("--2006-01-02 15:04:05--")
	fmt.Fprintf(LogOutput, "%s  %s\n", timestamp, normalizedURL)

	// Resolve hostname (host name ---> IPs) (using DNS lookup)
	// used tp print host ips
	if err := ResolveHostname(normalizedURL); err != nil {
		return err
	}

	// -P flag directories creation.
	InitializePath()

	response, err := MakeRequest(normalizedURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Handle the response
	filename, destPath, contentLength, err := HandleResponse(response)
	if err != nil {
		return err
	}

	// Now, download and save the file
	bytesDownloaded, err := SaveFile(response, destPath, contentLength, filename)
	if err != nil {
		return err
	}

	fmt.Fprintf(LogOutput, "\n%s (%s) saved [%d]\n", filepath.Base(destPath), normalizedURL, bytesDownloaded)
	return nil
}

// Resolves a hostname to its IP addresses using DNS lookup.
func ResolveHostname(URL string) error {
	// No need for error check as we already parsed the url before.
	parsedURL, _ := url.Parse(URL)

	hostname := parsedURL.Hostname()
	if hostname == "" {
		fmt.Fprintf(LogOutput, "%s: Invalid host name.\n", URL)
		return fmt.Errorf("Errooooor")
	}
	fmt.Fprintf(LogOutput, "Resolving %s (%s)... ", hostname, hostname)
	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Fprintf(LogOutput, "failed: %s.\n", err.Error())
		fmt.Fprintf(LogOutput, "wget: unable to resolve host address '%s'\n", hostname)
		return err
	}

	ipStrs := make([]string, len(ips))
	for i, ip := range ips {
		ipStrs[i] = ip.String()
	}
	fmt.Fprintf(LogOutput, "%s\n", strings.Join(ipStrs, ", "))
	return nil
}

// If -P flag full FilePath will have a string value (absolute path), here we create all the directories of that.
func InitializePath() {
	if FilePath == "" {
		return
	}
	err := os.MkdirAll(FilePath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
		os.Exit(1)
	}
}

// perform an HTTP request to the URL and return the response.
func MakeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Wget/1.21.1")
	
	resp, err := http.DefaultClient.Do(req)
	fmt.Fprintf(LogOutput, "HTTP request sent, awaiting response...  ")
	if err != nil {
		fmt.Fprintf(LogOutput, "Request failed: %v\n", err)
	}
	return resp, nil
}
