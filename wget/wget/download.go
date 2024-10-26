package wget

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

// SaveFile downloads the content from resp and saves it to destPath.
// It handles rate limiting and displays a progress bar if not in silent mode.
func SaveFile(resp *http.Response, destPath string, contentLength int64, filename string) (int64, error) {
	// Create the file
	file, err := os.Create(destPath)
	if err != nil {
		fmt.Fprintf(LogOutput, "wget: error creating file '%s': %v\n", destPath, err)
		return 0, err
	}
	defer file.Close()

	// Get rate limit in bytes per second
	rateLimit, err := GetRateLimit()
	if err != nil {
		fmt.Fprintf(LogOutput, "Error getting rate limit: %v\n", err)
		return 0, err
	}

	// Apply rate limiting if specified
	var reader io.Reader = resp.Body
	if rateLimit > 0 {
		reader = NewRateLimitedReader(resp.Body, rateLimit)
	}

	// Use progress bar if not silent
	if !Silent {
		p := mpb.New()
		var bar *mpb.Bar

		if contentLength > 0 {
			bar = p.AddBar(contentLength,
				mpb.BarStyle(" ▓▓░"),
				mpb.PrependDecorators(
					decor.Name(fmt.Sprintf(" %s ", filename)),
					decor.CountersKibiByte("%.1f / %.1f"),
				),
				mpb.AppendDecorators(
					decor.Percentage(),
					decor.AverageSpeed(decor.UnitKiB, " %.2f "),
					decor.EwmaETA(decor.ET_STYLE_GO, 60),
				),
			)
		} else {
			bar = p.AddBar(0,
				mpb.BarStyle(" ▓▓░"),
				mpb.PrependDecorators(
					decor.Name(fmt.Sprintf(" %s ", filename)),
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

		return bytesDownloaded, nil
	}
}

// New io.Reader type that wrap and limits the reading speed from an io.Reader.
type RateLimitedReader struct {
	reader    io.Reader
	rateLimit float64   // bytes per second
	lastTime  time.Time // time of the last read operation
	bytesRead int64     // total bytes read since lastTime
}

// Creates a new RateLimitedReader.
func NewRateLimitedReader(r io.Reader, rateLimit float64) *RateLimitedReader {
	return &RateLimitedReader{
		reader:    r,
		rateLimit: rateLimit,
		lastTime:  time.Now(),
	}
}

// Reads data from the underlying reader and enforces the rate limit.
func (r *RateLimitedReader) Read(p []byte) (int, error) {

	now := time.Now()
	elapsed := now.Sub(r.lastTime).Seconds()

	// Calculate the maximum bytes allowed to read based on elapsed time.
	allowedBytes := r.rateLimit * elapsed
    //fmt.Fprintln(LogOutput,"///////////////////////", allowedBytes, r.rateLimit, elapsed)
    //fmt.Fprintln(LogOutput,"++++++++++++++++++++++++++++++", r.bytesRead)
	// If we have read more than allowed, we need to wait.
	if float64(r.bytesRead) >= allowedBytes {
        //fmt.Fprintln(LogOutput,"----------------------------", r.bytesRead, allowedBytes)
		sleepTime := time.Duration(((float64(r.bytesRead) - allowedBytes) / r.rateLimit) * float64(time.Second))
		time.Sleep(sleepTime)

		now = time.Now()
		elapsed = now.Sub(r.lastTime).Seconds()
		allowedBytes = r.rateLimit * elapsed
	}

	// Calculate the remaining bytes we are allowed to read.
	remainingBytes := allowedBytes - float64(r.bytesRead)
	if remainingBytes < 1 {
		// Ensure we read at least 1 byte.
		remainingBytes = 1
	}

	// Limit the read to the remaining allowed bytes and buffer size.
	maxBytes := int(remainingBytes)
	if maxBytes > len(p) {
		maxBytes = len(p)
	}
	n, err := r.reader.Read(p[:maxBytes])
	r.bytesRead += int64(n)

	// Update lastTime and bytesRead if enough time has passed.
	if r.bytesRead >= int64(r.rateLimit) {       
		r.lastTime = now
		r.bytesRead = 0
	}

	return n, err
}

// GetRateLimit converts the rate limit to bytes per second.
func GetRateLimit() (float64, error) {
	var rateLimit float64
	switch strings.ToLower(RateLimitUnit) {
	case "b":
		rateLimit = RateLimit
	case "k":
		rateLimit = RateLimit * 1024
	case "m":
		rateLimit = RateLimit * 1024 * 1024
	default:
		rateLimit = RateLimit
	}
	return rateLimit, nil
}
