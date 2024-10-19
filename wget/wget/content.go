package wget

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/juju/ratelimit"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	"golang.org/x/sys/unix"
)

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
		bucket := ratelimit.NewBucketWithRate(float64(rateLimitBytesPerSec), 1) // Minimal capacity
		reader = ratelimit.Reader(resp.Body, bucket)
	}

	// Use progress bar if not Silent
	var bytesDownloaded int64
	if !Silent {
		p := mpb.New(mpb.WithWidth(int(float64(GetTerminalWidth()) * 0.7)))
		var bar *mpb.Bar

		if contentLength > 0 {
			bar = p.AddBar(contentLength,
				mpb.BarStyle(" ▓▓░"),
				mpb.PrependDecorators(
					decor.Name(fmt.Sprintf(" %s ", filename)),
					decor.CountersKibiByte("% .1f / % .1f"),
				),
				mpb.AppendDecorators(
					decor.Percentage(decor.WCSyncSpace),
					decor.AverageSpeed(decor.UnitKiB, " % .2f "),
					decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncWidth),
				),
			)
		} else {
			bar = p.AddBar(0,
				mpb.BarStyle(" ▓▓░"),
				mpb.PrependDecorators(
					decor.Name(fmt.Sprintf(" %s ", filename)),
				),
				mpb.AppendDecorators(
					decor.Percentage(),
					decor.AverageSpeed(decor.UnitKiB, " % .2f s"),
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

func GetRateLimitBytesPerSecond() (int64, error) {
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
	return int64(rateLimitBytesPerSec), nil
}

// Utility functions
func GetTerminalWidth() int {
	fd := int(os.Stdout.Fd())

	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return 80
	}

	return int(ws.Col)
}
