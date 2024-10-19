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
func DownloadFile(rawURL string) error {

	normalizedURL, err := HSTSurlCheck(rawURL)
	if err != nil {
		return err
	}

	timestamp := time.Now().Format("--2006-01-02 15:04:05--")
	fmt.Fprintf(LogOutput, "%s  %s\n", timestamp, normalizedURL)

	// No need for error check as we already parsed the url before.
	parsedURL, _ := url.Parse(normalizedURL)

	hostname := parsedURL.Hostname()
	if hostname == "" {
		fmt.Fprintf(LogOutput, "%s: Invalid host name.\n", normalizedURL)
		return fmt.Errorf("Errooooor")
	}

	// Resolve hostname (host name ---> IPs) (using DNS lookup)
	if err := ResolveHostname(hostname); err != nil {
		return err
	}

	// -P flag directories creation.
	InitializePath()

	// Prepare HTTP client
	//client := PrepareHTTPClient()

	// Prepare HTTP request
	//req, err := PrepareHTTPRequest(normalizedURL)
	// if err != nil {
	// 	return err
	// }

	// Perform the request.
	//resp, err := PerformRequest(req, client, normalizedURL)
	//resp, err := http.Get(normalizedURL)
	req, err := http.NewRequest("GET", normalizedURL, nil)
	req.Header.Set("User-Agent", "Wget/1.21.1")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Connection", "Keep-Alive")
	if err != nil {
		return err
	}
	client := http.Client{}
	fmt.Print("Sending request, awaiting response...  ")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(LogOutput, "Request failed: %v\n", err)
	}
	fmt.Println("--------------", resp.Header.Get("Content-Length"))
	defer resp.Body.Close()

	// Handle the response
	filename, destPath, contentLength, err := HandleResponse(resp, normalizedURL)
	if err != nil {
		return err
	}

	// Now, download the file
	bytesDownloaded, err := DownloadAndSaveFile(resp, destPath, contentLength, filename)
	if err != nil {
		return err
	}

	fmt.Fprintf(LogOutput, "\n%s (%s) saved [%d]\n", filepath.Base(destPath), hostname, bytesDownloaded)
	return nil
}

// Resolves a hostname to its IP addresses using DNS lookup.
func ResolveHostname(hostname string) error {
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
