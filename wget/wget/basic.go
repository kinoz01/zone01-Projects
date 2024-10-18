package wget

import (
    "crypto/tls"
    "fmt"
    "io"
    "mime"
    "net"
    "net/http"
    "net/url"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "sync"
    "time"

    "golang.org/x/exp/slices"
    "golang.org/x/sys/unix"

    "github.com/juju/ratelimit"
    "github.com/vbauerster/mpb"
    "github.com/vbauerster/mpb/decor"
)

func Run() {
    var wg sync.WaitGroup

    if Silent {
        logFileName := getLogFileName()
        logFile, err := os.Create(logFileName)
        if err != nil {
            fmt.Printf("Error creating log file '%s': %v\n", logFileName, err)
            os.Exit(1)
        }
        defer logFile.Close()
        LogOutput = logFile
    } else {
        LogOutput = os.Stdout
    }

    startTime := time.Now()
    fmt.Fprintf(LogOutput, "start at %s\n", startTime.Format("2006-01-02 15:04:05"))

    for _, rawURL := range URLs {
        wg.Add(1)
        go func(rawURL string) {
            defer wg.Done()
            err := downloadFile(rawURL)
            if err != nil {
                fmt.Fprintf(LogOutput, "Error downloading %s: %v\n", rawURL, err)
            } else {
                fmt.Fprintf(LogOutput, "\nDownloaded [%s]\n", rawURL)
            }
        }(rawURL)
    }

    wg.Wait()

    endTime := time.Now()
    fmt.Fprintf(LogOutput, "finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}

func downloadFile(rawURL string) error {
    // Normalize the URL
    normalizedURL, transformed, err := normalizeURL(rawURL)
    if err != nil {
        fmt.Fprintf(LogOutput, "Invalid URL '%s': %v\n", rawURL, err)
        return err
    }

    if transformed {
        fmt.Fprintf(LogOutput, "URL transformed to HTTPS due to an HSTS policy\n")
    }

    timestamp := time.Now().Format("--2006-01-02 15:04:05--")
    fmt.Fprintf(LogOutput, "%s  %s\n", timestamp, normalizedURL)

    // Parse the URL to extract the hostname
    parsedURL, err := url.Parse(normalizedURL)
    if err != nil {
        fmt.Fprintf(LogOutput, "wget: unable to parse URL '%s': %v\n", normalizedURL, err)
        return err
    }

    hostname := parsedURL.Hostname()
    if hostname == "" {
        fmt.Fprintf(LogOutput, "wget: unable to parse hostname from URL '%s'\n", normalizedURL)
        return fmt.Errorf("unable to parse hostname")
    }

    // Resolve hostname
    if err := resolveHostname(hostname); err != nil {
        return err
    }

    // Prepare HTTP client
    redirectCount := 0
    var resp *http.Response

    client := &http.Client{
        Timeout: 30 * time.Second,
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            if redirectCount >= 10 {
                return fmt.Errorf("maximum redirects reached")
            }
            redirectCount++
            fmt.Fprintf(LogOutput, "Location: %s [following]\n", req.URL.String())
            return nil
        },
        Transport: &http.Transport{
            TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
            ForceAttemptHTTP2: false,
        },
    }

    // Prepare HTTP request
    req, err := http.NewRequest("GET", normalizedURL, nil)
    if err != nil {
        fmt.Fprintf(LogOutput, "failed to create request: %v\n", err)
        return err
    }

    // Set headers to mimic wget
    req.Header.Set("User-Agent", "Wget/1.21.1")
    req.Header.Set("Accept", "*/*")
    req.Header.Set("Accept-Encoding", "identity")
    req.Header.Set("Connection", "Keep-Alive")

    fmt.Fprintf(LogOutput, "Connecting to %s (%s)... ", hostname, parsedURL.Host)
    resp, err = client.Do(req)
    if err != nil {
        fmt.Fprintf(LogOutput, "failed: %v\n", err)
        return err
    }
    defer resp.Body.Close()

    fmt.Fprintf(LogOutput, "connected.\n")
    fmt.Fprintf(LogOutput, "HTTP request sent, awaiting response... %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))

    if resp.StatusCode >= 300 && resp.StatusCode < 400 {
        location := resp.Header.Get("Location")
        if location != "" {
            fmt.Fprintf(LogOutput, "Location: %s [following]\n", location)
            return downloadFile(location)
        }
    }

    if resp.StatusCode != http.StatusOK {
        fmt.Fprintf(LogOutput, "wget: server returned error: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
        return fmt.Errorf("server returned error: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
    }

    // Now, determine the filename using the final URL after redirects
    filename, err := determineFilename(resp, OutputFile)
    if err != nil {
        return err
    }

    // Handle existing files
    destPath := getDestinationPath(filename)
    destPath = handleExistingFiles(destPath)

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
        sizeStr := formatSize(contentLength)
        fmt.Fprintf(LogOutput, "Length: %d (%s) [%s]\n", contentLength, sizeStr, contentType)
    } else {
        fmt.Fprintf(LogOutput, "Length: unspecified [%s]\n", contentType)
    }
    fmt.Fprintf(LogOutput, "Saving to: '%s'\n\n", destPath)

    // Now, download the file
    bytesDownloaded, err := downloadAndSaveFile(resp, destPath, contentLength, filename)
    if err != nil {
        return err
    }

    fmt.Fprintf(LogOutput, "\n%s (%s) saved [%d]\n", filepath.Base(destPath), hostname, bytesDownloaded)
    return nil
}

func resolveHostname(hostname string) error {
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

func normalizeURL(rawURL string) (string, bool, error) {
    transformed := false

    rawURL = strings.TrimSpace(rawURL)

    if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
        rawURL = "http://" + rawURL
    }
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        return "", false, err
    }

    client := &http.Client{
        Timeout: 10 * time.Second,
        CheckRedirect: func(res *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }

    request, err := http.NewRequest("HEAD", parsedURL.String(), nil)
    if err != nil {
        return "", false, err
    }

    resp, err := client.Do(request)
    if err == nil {
        if resp.StatusCode >= 300 && resp.StatusCode < 400 {
            location := resp.Header.Get("Location")
            if strings.HasPrefix(location, "https://") {
                parsedURL.Scheme = "https"
                transformed = true
            }
        }
        if hsts := resp.Header.Get("Strict-Transport-Security"); hsts != "" {
            parsedURL.Scheme = "https"
            transformed = true
        }
        resp.Body.Close()
    }
    normalizedURL := parsedURL.String()
    return normalizedURL, transformed, nil
}

func determineFilename(resp *http.Response, outputFile string) (string, error) {
    if outputFile != "" {
        return outputFile, nil
    }

    cd := resp.Header.Get("Content-Disposition")
    if cd != "" {
        _, params, err := mime.ParseMediaType(cd)
        if err == nil {
            if filename, ok := params["filename"]; ok {
                if isInvalidFilename(filename) {
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

    if isInvalidFilename(filename) {
        filename = "index.html"
    } else if !strings.Contains(filename, ".") {
        contentType := resp.Header.Get("Content-Type")
        ext := getFileExtension(contentType)
        filename += ext
    }

    return filename, nil
}

func isInvalidFilename(filename string) bool {
    invalidFilenames := []string{"", ".", "/", "unsupportedbrowser"}
    if slices.Contains(invalidFilenames, filename) || strings.HasSuffix(filename, "/") || strings.Contains(filename, "?") {
        return true
    }
    return false
}

func getFileExtension(contentType string) string {
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

func getDestinationPath(filename string) string {
    destDir := FilePath
    return filepath.Join(destDir, filename)
}

func handleExistingFiles(destPath string) string {
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

func formatSize(size int64) string {
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

func getRateLimitBytesPerSecond() (int64, error) {
    if RateLimitUnit == "" || RateLimit == 0 {
        return 0, nil // No rate limit
    }
    var rateLimitBytesPerSec float64
    switch strings.ToLower(RateLimitUnit) {
    case "b":
        rateLimitBytesPerSec = RateLimit
    case "k":
        rateLimitBytesPerSec = RateLimit * 1024
    case "m":
        rateLimitBytesPerSec = RateLimit * 1024 * 1024
    case "g":
        rateLimitBytesPerSec = RateLimit * 1024 * 1024 * 1024
    default:
        return 0, fmt.Errorf("invalid RateLimitUnit: %s", RateLimitUnit)
    }
    return int64(rateLimitBytesPerSec), nil
}

func getLogFileName() string {
    baseName := "wget-log"
    fileName := baseName
    i := 1
    for {
        if _, err := os.Stat(fileName); os.IsNotExist(err) {
            return fileName
        }
        fileName = fmt.Sprintf("%s.%d", baseName, i)
        i++
    }
}

func downloadAndSaveFile(resp *http.Response, destPath string, contentLength int64, filename string) (int64, error) {
    // Create the file
    file, err := os.Create(destPath)
    if err != nil {
        fmt.Fprintf(LogOutput, "wget: error creating file '%s': %v\n", destPath, err)
        return 0, err
    }
    defer file.Close()

    // Get rate limit in bytes per second
    rateLimitBytesPerSec, err := getRateLimitBytesPerSecond()
    if err != nil {
        fmt.Fprintf(LogOutput, "error with rate limit: %v\n", err)
        return 0, err
    }

    // Apply rate limiting if specified
    var reader io.Reader = resp.Body
    if rateLimitBytesPerSec > 0 {
        bucket := ratelimit.NewBucketWithRate(float64(rateLimitBytesPerSec), rateLimitBytesPerSec)
        reader = ratelimit.Reader(resp.Body, bucket)
    }

    // Use progress bar if not Silent
    var bytesDownloaded int64
    if !Silent {
        p := mpb.New(mpb.WithWidth(int(float64(getTerminalWidth()) * 0.7)))
        var bar *mpb.Bar

        if contentLength > 0 {
            bar = p.AddBar(contentLength,
                mpb.BarStyle(" ▓▓░"),
                mpb.PrependDecorators(
                    decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncWidth), // Time remaining on the left
                    decor.Name(fmt.Sprintf(" %s ", filename)),
                    decor.CountersKibiByte("% .1f / % .1f"),
                ),
                mpb.AppendDecorators(
                    decor.Percentage(decor.WCSyncSpace),
                    decor.AverageSpeed(decor.UnitKiB, " % .2f KiB/s"),
                ),
            )
        } else {
            bar = p.AddBar(0,
                mpb.BarStyle(" ▓▓░"),
                mpb.PrependDecorators(
                    decor.Name(fmt.Sprintf(" %s ", filename)),
                ),
                mpb.AppendDecorators(
                    decor.AverageSpeed(decor.UnitKiB, " % .2f KiB/s"),
                    decor.Percentage(),
                ),
            )
        }

        reader = bar.ProxyReader(reader)

        bytesDownloaded, err = io.Copy(file, reader)
        if err != nil {
            fmt.Fprintf(LogOutput, "wget: error writing to file '%s': %v\n", destPath, err)
            return bytesDownloaded, err
        }

        if contentLength <= 0 {
            bar.SetTotal(bar.Current(), true)
        }

        p.Wait()
    } else {
        bytesDownloaded, err = io.Copy(file, reader)
        if err != nil {
            fmt.Fprintf(LogOutput, "wget: error writing to file '%s': %v\n", destPath, err)
            return bytesDownloaded, err
        }
    }

    return bytesDownloaded, nil
}

// Utility functions
func getTerminalWidth() int {
    fd := int(os.Stdout.Fd())

    ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
    if err != nil {
        return 80
    }

    return int(ws.Col)
}