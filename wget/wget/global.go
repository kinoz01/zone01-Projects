package wget

import (
	"io"

	"github.com/vbauerster/mpb/v8"
)

var (
	NotFlags         []string // All non flags found.
	Silent           bool  // -B flag (write to wget-log)
	RateLimit        float64
	RateLimitUnit    string
	OutputFile       string
	FilePath         string
	URLs             []string // List of URLs to download
	InputError1      error
	InputError2      bool
	Mirror           bool
	RejectedSuffixes []string
	ExcludedPaths    []string
	ConvertLinks     bool
	ProgressBar      *mpb.Progress
	LogOutput        io.Writer
)
