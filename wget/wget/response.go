package wget

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"golang.org/x/exp/slices"
)

// Processes the HTTP response, handles redirects, and prepares for download.
func HandleResponse(resp *http.Response,  urlStr string) (string, string, int64, error) {
	// if resp.StatusCode >= 300 && resp.StatusCode < 400 {
	// 	location := resp.Header.Get("Location")
	// 	if location != "" {
	// 		fmt.Fprintf(LogOutput, "Location: %s [following]\n", location)
	// 		// Prepare new request
	// 		req, err := PrepareHTTPRequest(location)
	// 		if err != nil {
	// 			return "", "", 0, err
	// 		}
	// 		// Perform the request recursively
	// 		newResp, err := PerformRequest(req, client, location)
	// 		if err != nil {
	// 			return "", "", 0, err
	// 		}
	// 		defer newResp.Body.Close()
	// 		return HandleResponse(newResp, client, location)
	// 	}
	// }

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(LogOutput, "\nwget: server returned error: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		return "", "", 0, fmt.Errorf("Erooooor")
	} else {
		fmt.Fprintf(LogOutput, "%d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Determine the filename
	filename, err := DetermineFilename(resp)
	if err != nil {
		return "", "", 0, err
	}
	OutputFile = filename

	// Handle existing files
	destPath := HandleExistingFiles()

	// Get content length and content type
	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")
	if contentLength == -1 {		
		if cl := resp.Header.Get("Content-Length"); cl != "" {
			contentLength, _ = strconv.ParseInt(cl, 10, 64)
		}
	}

	// Print length and saving message
	if contentLength > 0 {
		sizeStr := FormatSize(contentLength)
		fmt.Fprintf(LogOutput, "Length: %d (%s) [%s]\n", contentLength, sizeStr, contentType)
	} else {
		fmt.Fprintf(LogOutput, "Length: unspecified [%s]\n", contentType)
	}
	fmt.Fprintf(LogOutput, "Saving to: '%s'\n\n", destPath)

	return filename, destPath, contentLength, nil
}

// Determines the filename for saving the downloaded content
func DetermineFilename(resp *http.Response) (string, error) {
	if OutputFile != "" {
		return OutputFile, nil
	}

	cd := resp.Header.Get("Content-Disposition")
	if cd != "" {
		_, params, err := mime.ParseMediaType(cd)
		if err == nil {
			if filename, ok := params["filename"]; ok {
				if IsInvalidFilename(filename) {
					return "index.html", nil
				}
				return filename, nil
			}
		}
	}

	// Use the last path segment of the final URL
	finalURL := resp.Request.URL
	path := finalURL.Path
	filename := filepath.Base(path)

	if IsInvalidFilename(filename) {
		filename = "index.html"
	} else if !strings.Contains(filename, ".") {
		contentType := resp.Header.Get("Content-Type")
		ext := GetFileExtension(contentType)
		filename += ext
	}

	return filename, nil
}

// Checks if the destination file exists and modifies the name if necessary
func HandleExistingFiles() string {
	destDir := FilePath
	destPath := filepath.Join(destDir, OutputFile)

	dir := filepath.Dir(destPath)
	ext := filepath.Ext(destPath)
	name := strings.TrimSuffix(filepath.Base(destPath), ext)

	i := 1
	originalDestPath := destPath
	for {
		if _, err := os.Stat(destPath); os.IsNotExist(err) {
			break
		}
		destPath = filepath.Join(dir, fmt.Sprintf("%s.%d%s", name, i, ext))
		i++
	}
	if destPath != originalDestPath {
		fmt.Fprintf(LogOutput, "File '%s' already exists, saving as '%s'\n", originalDestPath, destPath)
	}
	return destPath
}

func IsInvalidFilename(filename string) bool {
	invalidFilenames := []string{"", ".", "/", "unsupportedbrowser"}
	if slices.Contains(invalidFilenames, filename) || strings.HasSuffix(filename, "/") || strings.Contains(filename, "?") {
		return true
	}
	return false
}

func GetFileExtension(contentType string) string {
	if contentType == "" {
		return ""
	}
	if strings.Contains(contentType, ";") {
		contentType = strings.Split(contentType, ";")[0]
	}
	switch contentType {
	case "text/html":
		return ".html"
	case "text/plain":
		return ".txt"
	case "application/json":
		return ".json"
	case "application/octet-stream":
		return ".bin"
	default:
		exts, _ := mime.ExtensionsByType(contentType)
		if len(exts) > 0 {
			return exts[0]
		}
	}
	return ""
}

func FormatSize(size int64) string {
	const (
		KB = 1 << (10 * 1)
		MB = 1 << (10 * 2)
		GB = 1 << (10 * 3)
	)

	floatSize := float64(size)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", floatSize/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", floatSize/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", floatSize/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}
