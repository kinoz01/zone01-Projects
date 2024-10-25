package wget

import (
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "sync"
    "time"

    "golang.org/x/net/html"
)

// MirrorWebsite mirrors a website by downloading its content and resources.
func MirrorWebsite(targetURL string, rejectList, excludeList []string, convertLinks bool, log *os.File, wg *sync.WaitGroup) {
    defer wg.Done()

    startTime := time.Now().Format("2006-01-02 15:04:05")
    fmt.Fprintf(log, "Start at %s\n", startTime)

    resp, err := http.Get(targetURL)
    if err != nil {
        fmt.Fprintf(log, "Error: %v\n", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Fprintf(log, "Status %s\n", resp.Status)
        return
    }

    docBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Fprintf(log, "Error reading response body: %v\n", err)
        return
    }

    baseURL := strings.TrimSuffix(targetURL, "/")
    outputDir := filepath.Join(".", sanitizeFilename(baseURL))

    excludeList = cleanList(excludeList)
    if shouldExclude(outputDir, excludeList) {
        fmt.Fprintf(log, "Skipping directory %s due to exclude list\n", outputDir)
        return
    }

    err = os.MkdirAll(outputDir, os.ModePerm)
    if err != nil {
        fmt.Fprintf(log, "Error creating directory %s: %v\n", outputDir, err)
        return
    }

    outputFilePath := filepath.Join(outputDir, "index.html")

    node, err := html.Parse(strings.NewReader(string(docBytes)))
    if err != nil {
        fmt.Fprintf(log, "Error parsing HTML: %v\n", err)
        return
    }

    var wgResources sync.WaitGroup

    processNode(node, baseURL, outputDir, rejectList, excludeList, convertLinks, log, &wgResources)

    wgResources.Wait()

    var outputHTML strings.Builder
    err = html.Render(&outputHTML, node)
    if err != nil {
        fmt.Fprintf(log, "Error rendering HTML: %v\n", err)
        return
    }

    err = os.WriteFile(outputFilePath, []byte(outputHTML.String()), 0644)
    if err != nil {
        fmt.Fprintf(log, "Error writing to file %s: %v\n", outputFilePath, err)
        return
    }

    fmt.Fprintf(log, "Mirrored %s to %s\n", targetURL, outputFilePath)
    fmt.Fprintf(log, "Finished at %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

// processNode processes an HTML node and its children, handling resources and converting links if needed.
func processNode(n *html.Node, baseURL, outputDir string, rejectList, excludeList []string, convertLinks bool, log *os.File, wg *sync.WaitGroup) {
    if n.Type == html.ElementNode {
        var attrKey, attrValue string

        switch n.Data {
        case "a", "link", "img", "script":
            // Determine the attribute to process based on the tag
            attrKey = getAttributeKey(n.Data)
            if attrKey != "" {
                for i, a := range n.Attr {
                    if a.Key == attrKey {
                        attrValue = a.Val
                        // Handle the resource
                        newVal := handleResource(attrValue, baseURL, outputDir, rejectList, excludeList, convertLinks, log, wg)
                        if convertLinks && newVal != "" {
                            n.Attr[i].Val = newVal
                        }
                        break
                    }
                }
            }
        case "style":
            // Process inline CSS
            if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
                cssContent := n.FirstChild.Data
                newCSSContent := processCSSContent(cssContent, baseURL, outputDir, rejectList, excludeList, convertLinks, log, wg)
                if convertLinks {
                    n.FirstChild.Data = newCSSContent
                }
            }
        }

        // Process style attribute
        for i, a := range n.Attr {
            if a.Key == "style" {
                cssContent := a.Val
                newCSSContent := processCSSContent(cssContent, baseURL, outputDir, rejectList, excludeList, convertLinks, log, wg)
                if convertLinks {
                    n.Attr[i].Val = newCSSContent
                }
                break
            }
        }
    }

    // Recursively process child nodes
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        processNode(c, baseURL, outputDir, rejectList, excludeList, convertLinks, log, wg)
    }
}

// getAttributeKey returns the attribute key to process based on the tag name.
func getAttributeKey(tagName string) string {
    switch tagName {
    case "a", "link":
        return "href"
    case "img", "script":
        return "src"
    default:
        return ""
    }
}

// handleResource handles downloading a resource and optionally converting the link.
func handleResource(attrValue, baseURL, outputDir string, rejectList, excludeList []string, convertLinks bool, log *os.File, wg *sync.WaitGroup) string {
    // Skip empty attributes
    if attrValue == "" {
        return ""
    }

    absoluteURL := resolveURL(baseURL, attrValue)

    if shouldReject(absoluteURL, rejectList) {
        fmt.Fprintf(log, "Skipping download of %s due to reject list\n", absoluteURL)
        return ""
    }

    relativePath := getRelativePath(absoluteURL, baseURL)
    outputPath := filepath.Join(outputDir, relativePath)

    if shouldExclude(outputPath, excludeList) {
        fmt.Fprintf(log, "Skipping %s due to exclude list\n", outputPath)
        return ""
    }

    wg.Add(1)
    go func() {
        defer wg.Done()
        err := downloadResource(absoluteURL, outputPath, log)
        if err != nil {
            fmt.Fprintf(log, "Error downloading resource %s: %v\n", absoluteURL, err)
        }
    }()

    if convertLinks {
        // Return the new attribute value pointing to the local file
        return filepath.ToSlash(relativePath)
    }
    return ""
}

// processCSSContent processes CSS content, downloading resources and converting URLs if needed.
func processCSSContent(cssContent, baseURL, outputDir string, rejectList, excludeList []string, convertLinks bool, log *os.File, wg *sync.WaitGroup) string {
    urls := extractURLsFromCSS(cssContent)
    for _, urlStr := range urls {
        absoluteURL := resolveURL(baseURL, urlStr)

        if shouldReject(absoluteURL, rejectList) {
            fmt.Fprintf(log, "Skipping CSS resource %s due to reject list\n", absoluteURL)
            continue
        }

        relativePath := getRelativePath(absoluteURL, baseURL)
        outputPath := filepath.Join(outputDir, relativePath)

        if shouldExclude(outputPath, excludeList) {
            fmt.Fprintf(log, "Skipping %s due to exclude list\n", outputPath)
            continue
        }

        wg.Add(1)
        go func(url, path string) {
            defer wg.Done()
            err := downloadResource(url, path, log)
            if err != nil {
                fmt.Fprintf(log, "Error downloading CSS resource %s: %v\n", url, err)
            }
        }(absoluteURL, outputPath)

        if convertLinks {
            // Update the CSS content to point to the local file
            localURL := filepath.ToSlash(relativePath)
            cssContent = strings.ReplaceAll(cssContent, urlStr, localURL)
        }
    }
    return cssContent
}

// downloadResource downloads a resource from the given URL and saves it to the specified output path.
func downloadResource(resourceURL, outputPath string, log *os.File) error {
    // Create necessary directories
    err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating directories for %s: %v", outputPath, err)
    }

    // Skip if file already exists
    if _, err := os.Stat(outputPath); err == nil {
        fmt.Fprintf(log, "Resource %s already downloaded\n", outputPath)
        return nil
    }

    resp, err := http.Get(resourceURL)
    if err != nil {
        return fmt.Errorf("error fetching resource %s: %v", resourceURL, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("error downloading resource %s: %s", resourceURL, resp.Status)
    }

    file, err := os.Create(outputPath)
    if err != nil {
        return fmt.Errorf("error creating file %s: %v", outputPath, err)
    }
    defer file.Close()

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        return fmt.Errorf("error writing to file %s: %v", outputPath, err)
    }

    fmt.Fprintf(log, "Downloaded %s to %s\n", resourceURL, outputPath)
    return nil
}

// resolveURL resolves a relative URL against the base URL.
func resolveURL(baseURL, link string) string {
    base, err := url.Parse(baseURL)
    if err != nil {
        return link
    }
    href, err := url.Parse(link)
    if err != nil {
        return link
    }
    return base.ResolveReference(href).String()
}

// getRelativePath computes the relative path of an absolute URL with respect to the base URL.
func getRelativePath(absoluteURL, baseURL string) string {
    base, err := url.Parse(baseURL)
    if err != nil {
        return ""
    }
    abs, err := url.Parse(absoluteURL)
    if err != nil {
        return ""
    }
    relPath, err := filepath.Rel(base.Path, abs.Path)
    if err != nil {
        return strings.TrimLeft(abs.Path, "/")
    }
    return relPath
}

// shouldReject determines if a resource URL should be rejected based on the reject list.
func shouldReject(resourceURL string, rejectList []string) bool {
    for _, pattern := range rejectList {
        if strings.HasSuffix(resourceURL, pattern) {
            return true
        }
    }
    return false
}

// shouldExclude determines if a path should be excluded based on the exclude list.
func shouldExclude(path string, excludeList []string) bool {
    for _, pattern := range excludeList {
        if strings.Contains(path, pattern) {
            return true
        }
    }
    return false
}

// cleanList cleans the input list by returning nil if it's empty or contains only empty strings.
func cleanList(list []string) []string {
    if len(list) == 0 || (len(list) == 1 && list[0] == "") {
        return nil
    }
    return list
}

// extractURLsFromCSS extracts URLs from CSS content.
func extractURLsFromCSS(cssContent string) []string {
    var urls []string
    cssURLRegex := regexp.MustCompile(`url\(['"]?([^'")]+)['"]?\)`)
    matches := cssURLRegex.FindAllStringSubmatch(cssContent, -1)
    for _, match := range matches {
        if len(match) > 1 {
            urls = append(urls, match[1])
        }
    }
    return urls
}

// sanitizeFilename sanitizes a URL to create a safe directory name.
func sanitizeFilename(urlStr string) string {
    u, err := url.Parse(urlStr)
    if err != nil {
        return strings.ReplaceAll(urlStr, "/", "_")
    }
    host := strings.ReplaceAll(u.Host, ".", "_")
    path := strings.ReplaceAll(u.Path, "/", "_")
    return host + path
}
