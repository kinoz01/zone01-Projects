package wget

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Parse input arguments, check errors and initialize global variables
func ParseArgs(args []string) {
	CheckFlags(args)
	CheckFlagsWithoutValue(args)
	GetFlagsValues(args)
	CheckErrors()
}

// Return if we find invalid flag argument or we find a help flag.
func CheckFlags(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if NotFlag(arg) {
				NotFlags = append(NotFlags, arg)
			}
		}
		if (arg == "-h" || arg == "--help") && len(NotFlags) == 0 {
			PrintHelp()
		}
	}
	if len(NotFlags) != 0 {
		PrintNotFlag()
	}
}

// Check if an argument that starts with "-" is a flag or not.
func NotFlag(arg string) bool {
	reh := regexp.MustCompile(`\A-h$`)
	reHelp := regexp.MustCompile(`\A--help$`)
	reB := regexp.MustCompile(`\A-B$`)
	reO := regexp.MustCompile(`\A-O($|=).*`)
	reP := regexp.MustCompile(`\A-P($|=).*`)
	reRate := regexp.MustCompile(`\A--rate-limit($|=).*`)
	rei := regexp.MustCompile(`\A-i($|=).*`)
	reMirror := regexp.MustCompile(`\A--mirror$`)
	reReject := regexp.MustCompile(`\A--reject($|=).*`)
	reR := regexp.MustCompile(`\A-R($|=).*`)
	reExclude := regexp.MustCompile(`\A--exclude($|=).*`)
	reX := regexp.MustCompile(`\A-X($|=).*`)
	reConvert := regexp.MustCompile(`\A--convert-links$`)
	if reh.MatchString(arg) || reHelp.MatchString(arg) || reB.MatchString(arg) ||
		reO.MatchString(arg) || reP.MatchString(arg) || reRate.MatchString(arg) ||
		rei.MatchString(arg) || reMirror.MatchString(arg) || reReject.MatchString(arg) ||
		reR.MatchString(arg) || reExclude.MatchString(arg) || reX.MatchString(arg) || reConvert.MatchString(arg) {
		return false
	}
	return true
}

// Check if all arguments that need a value have one.
func CheckFlagsWithoutValue(args []string) {
	for _, arg := range args {
		if FlagWithoutVal(arg) {
			PrintFlagWithoutVal(arg)
		}
	}
}

// Return true if a flag don't have value.
func FlagWithoutVal(arg string) bool {
	reO := regexp.MustCompile(`\A-O(=|$)$`)
	reP := regexp.MustCompile(`\A-P(=|$)$`)
	reRate := regexp.MustCompile(`\A--rate-limit(=|$)$`)
	rei := regexp.MustCompile(`\A-i(=|$)$`)
	reReject := regexp.MustCompile(`\A--reject(=|$)$`)
	reR := regexp.MustCompile(`\A-R(=|$)$`)
	reExclude := regexp.MustCompile(`\A--exclude(=|$)$`)
	reX := regexp.MustCompile(`\A-X(=|$)$`)

	if reO.MatchString(arg) || reP.MatchString(arg) || reRate.MatchString(arg) ||
		rei.MatchString(arg) || reReject.MatchString(arg) || reR.MatchString(arg) ||
		reExclude.MatchString(arg) || reX.MatchString(arg) {
		return true // Match found
	}
	return false
}

// Initialise global variable (flags) values.
func GetFlagsValues(args []string) {
	for _, arg := range args {
		if arg == "-B" {
			Silent = true
		} else if strings.HasPrefix(arg, "-O=") {
			OutputFile = strings.TrimPrefix(arg, "-O=")
		} else if strings.HasPrefix(arg, "-P=") {
			HandlePath(strings.TrimPrefix(arg, "-P="))
		} else if strings.HasPrefix(arg, "--rate") {
			HandleRateLimit(arg)
		} else if strings.HasPrefix(arg, "-i=") {
			HandleInputFile(strings.TrimPrefix(arg, "-i="))
		} else if arg == "--mirror" {
			Mirror = true
		} else if strings.HasPrefix(arg, "--reject") {
			RejectedSuffixes = append(RejectedSuffixes, strings.Split(strings.TrimPrefix(arg, "--reject="), ",")...)
		} else if strings.HasPrefix(arg, "-R") {
			RejectedSuffixes = append(RejectedSuffixes, strings.Split(strings.TrimPrefix(arg, "-R="), ",")...)
		} else if strings.HasPrefix(arg, "--exclude") {
			ExcludedPaths = append(ExcludedPaths, strings.Split(strings.TrimPrefix(arg, "--exclude="), ",")...)
		} else if strings.HasPrefix(arg, "-X") {
			ExcludedPaths = append(ExcludedPaths, strings.Split(strings.TrimPrefix(arg, "-X="), ",")...)
		} else if arg == "--convert-links" {
			ConvertLinks = true
		} else {
			URLs = append(URLs, arg)
		}
	}
}

// Handle rate limit by returning if error and else initialising rate limit value.
func HandleRateLimit(arg string) {
	reRate := regexp.MustCompile(`\A--rate-limit=(\d+(\.\d+)?)([BMk]*)$`)
	if !reRate.MatchString(arg) {
		reWrongRate := regexp.MustCompile(`\A--rate-limit=(.*)$`)
		PrintInvalidRate(reWrongRate.FindStringSubmatch(arg)[1])
	} else {
		var err error
		RateLimit, err = strconv.ParseFloat(reRate.FindStringSubmatch(arg)[1], 64)
		if err != nil || RateLimit == 0 {
			PrintInvalidRate(reRate.FindStringSubmatch(arg)[1])
		}
		RateLimitUnit = reRate.FindStringSubmatch(arg)[3]
	}
}

// Initialise the FilePath with an absolute path.
func HandlePath(path string) {
	var err error
	path, err = MakePath(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	FilePath = path
}

// Store lines of input file if it exist in the URLs array.
// If file don't exist InputError1 != nil
// If file is empty InputError2 = true
func HandleInputFile(path string) {
	var err1 error
	noInputs := true
	path, err1 = MakePath(path)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err1)
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		InputError1 = err
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		URLs = append(URLs, line)
		noInputs = false
	}
	if noInputs {
		InputError2 = true
	}
}

// Expands '~' to the user's home directory and converts
// relative paths to absolute paths. It returns the processed path.
func MakePath(path string) (string, error) {

	// Expand '~' to the user's home directory
	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("error getting current user: %v", err)
		}
		path = filepath.Join(usr.HomeDir, strings.TrimPrefix(path, "~"))
	}

	// Convert to absolute path if it's relative.
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return "", fmt.Errorf("error getting absolute path: %v", err)
		}
		path = absPath
	}
	return path, nil
}

// Check the URLs slice isn't empty and initialize Warnings if found.
func CheckErrors() {
	if len(URLs) == 0 {
		PrintMissingURL()
	}
	if Mirror {
		if OutputFile != "" {
			Warnings += "WARNING: You can't change filenames when mirroring a website!\n"
		}
		if FilePath != "" {
			Warnings += "WARNING: You can't change download path when mirroring a website!\n"
		}
		if RateLimit != 0 {
			Warnings += "WARNING: You can't use rate limit when mirroring a website!\n"
		}
		if len(URLs) > 1 {
			Warnings += "WARNING: Mirroring support only one link at a time! I mirrored the first link in your list\n"
		}
	} else {
		if len(RejectedSuffixes) != 0 {
			Warnings += "WARNING: You can't use --reject or -R flags without using --mirror!\n"
		}
		if len(ExcludedPaths) != 0 {
			Warnings += "WARNING: You can't use --exclude or -X flags without using --mirror!\n"
		}
		if ConvertLinks {
			Warnings += "WARNING: You can't use --convert-links flags without using --mirror!\n"
		}
	}
}
