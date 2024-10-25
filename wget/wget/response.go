package wget

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

// Processes the HTTP response, and prepares for download.
func HandleResponse(resp *http.Response) (string, string, int64, error) {

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

	// Use the last path segment of the final URL
	filename := filepath.Base(resp.Request.URL.Path)

	if IsInvalidFilename(filename) {
		filename = "index.html"
	}
	return filename, nil
}

// Checks if the destination file exists and modifies the name if necessary
func HandleExistingFiles() string {

	destPath := filepath.Join(FilePath, OutputFile)

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
	return slices.Contains(invalidFilenames, filename)
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
