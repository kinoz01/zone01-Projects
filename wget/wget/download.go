package wget

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "time"

    "golang.org/x/sys/unix"

    "github.com/vbauerster/mpb"
    "github.com/vbauerster/mpb/decor"
)

// DownloadAndSaveFile downloads the content from resp and saves it to destPath.
// It handles rate limiting and displays a progress bar if not in silent mode.
func DownloadAndSaveFile(resp *http.Response, destPath string, contentLength int64, filename string) (int64, error) {
    // Create the file
    file, err := os.Create(destPath)
    if err != nil {
        fmt.Fprintf(LogOutput, "wget: error creating file '%s': %v\n", destPath, err)
        return 0, err
    }
    defer file.Close()

    // Get rate limit in bytes per second
    rateLimitBytesPerSec, err := GetRateLimitBytesPerSecond()
    if err != nil {
        fmt.Fprintf(LogOutput, "error with rate limit: %v\n", err)
        return 0, err
    }

    // Apply rate limiting if specified
    var reader io.Reader = resp.Body
    if rateLimitBytesPerSec > 0 {
        reader = NewRateLimitedReader(resp.Body, rateLimitBytesPerSec)
    }

    // Use progress bar if not silent
    if !Silent {
        p := mpb.New(mpb.WithWidth(int(float64(GetTerminalWidth()) * 0.7)))
        var bar *mpb.Bar

        if contentLength > 0 {
            bar = p.AddBar(contentLength,
                mpb.BarStyle(" ▓▓░"),
                mpb.PrependDecorators(
                    decor.Name(fmt.Sprintf(" %s ", filename)),
                    decor.CountersKibiByte("%.1f / %.1f"),
                ),
                mpb.AppendDecorators(
                    decor.Percentage(decor.WCSyncSpace),
                    decor.AverageSpeed(decor.UnitKiB, " %.2f "),
                    decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncWidth),
                ),
            )
        } else {
            bar = p.AddBar(0,
                mpb.BarStyle(" ▓▓░"),
                mpb.PrependDecorators(
                    decor.Name(fmt.Sprintf(" %s ", filename)),
                    //decor.CounterKibiByte("%.1f"),
					//decor.CountersKibiByte("% .1f"),
					decor.CountersKibiByte("%.1f / %.1f"),
                ),
                mpb.AppendDecorators(
                    decor.AverageSpeed(decor.UnitKiB, " %.2f "),
                ),
            )
        }

        reader = bar.ProxyReader(reader)
        bytesDownloaded, err := io.Copy(file, reader)
        if err != nil {
            fmt.Fprintf(LogOutput, "wget: error writing to file '%s': %v\n", destPath, err)
            return bytesDownloaded, err
        }

        // For unknown content length, set bar total to bytes downloaded and complete the bar
        if contentLength <= 0 {
            bar.SetTotal(bar.Current(), true)
        }

        p.Wait()

        return bytesDownloaded, nil
    } else {
        bytesDownloaded, err := io.Copy(file, reader)
        if err != nil {
            fmt.Fprintf(LogOutput, "wget: error writing to file '%s': %v\n", destPath, err)
            return bytesDownloaded, err
        }

        fmt.Fprintf(LogOutput, "Downloaded %d bytes\n", bytesDownloaded)

        return bytesDownloaded, nil
    }
}

// RateLimitedReader limits the reading speed from an io.Reader.
type RateLimitedReader struct {
    reader    io.Reader
    rateLimit float64    // bytes per second
    lastTime  time.Time  // last time bytes were read
    allowance float64    // bytes allowed to read
}

// NewRateLimitedReader creates a new RateLimitedReader.
func NewRateLimitedReader(r io.Reader, rateLimit float64) *RateLimitedReader {
    return &RateLimitedReader{
        reader:    r,
        rateLimit: rateLimit,
        lastTime:  time.Now(),
        allowance: rateLimit, // Start with full allowance
    }
}

// Read reads data from the underlying reader and enforces the rate limit.
func (r *RateLimitedReader) Read(p []byte) (int, error) {
    if r.rateLimit <= 0 {
        return r.reader.Read(p)
    }

    now := time.Now()
    elapsed := now.Sub(r.lastTime).Seconds()
    r.lastTime = now
    r.allowance += elapsed * r.rateLimit
    if r.allowance > r.rateLimit {
        r.allowance = r.rateLimit
    }

    maxBytes := int(r.allowance)
    if maxBytes < 1 {
        // Need to wait
        sleepTime := time.Duration((1 - r.allowance/r.rateLimit) * float64(time.Second))
        time.Sleep(sleepTime)
        r.allowance = 0
        maxBytes = 1
    }

    if maxBytes > len(p) {
        maxBytes = len(p)
    }

    n, err := r.reader.Read(p[:maxBytes])
    r.allowance -= float64(n)
    return n, err
}

// GetRateLimitBytesPerSecond converts the rate limit to bytes per second.
func GetRateLimitBytesPerSecond() (float64, error) {
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
    default:
        rateLimitBytesPerSec = RateLimit
    }
    return rateLimitBytesPerSec, nil
}

// GetTerminalWidth gets the width of the terminal for progress bar display.
func GetTerminalWidth() int {
    fd := int(os.Stdout.Fd())

    ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
    if err != nil {
        return 80
    }

    return int(ws.Col)
}
