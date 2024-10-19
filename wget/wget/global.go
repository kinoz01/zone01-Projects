package wget

import (
	"io"

	"github.com/vbauerster/mpb/v8"
)

var (
	NotFlags         []string      // All non flags found.
	Silent           bool          // -B flag (write to wget-log)
	RateLimit        float64       // Set the speed of download
	RateLimitUnit    string        // Unit of the rate (B, k or M)
	OutputFile       string        // Downloaded file name
	FilePath         string        // Saving path
	URLs             []string      // List of URLs to download
	InputError1      error         // File not found
	InputError2      bool          // File is found but its empty
	InputFileName    string        // To print file name in case of error
	Mirror           bool          // --mirror
	RejectedSuffixes []string      // rejected suffixes
	ExcludedPaths    []string      // rejected paths
	ConvertLinks     bool          // --convert-links
	ProgressBar      *mpb.Progress // Progress bar
	LogOutput        io.Writer     // writer to log (wget-file or stdout)
	Warnings         string        // Warnings found during CheckErrors() func
)
