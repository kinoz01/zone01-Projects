package wget

import (
	"fmt"
	"os"
	"strings"
)

// Printed when we can't find url to be resolved.
func PrintMissingURL() {
	fmt.Println("wget: missing URL\nUsage: go run . [OPTION]... [URL]...\n\nTry `go run  . --help' for more options.")
	os.Exit(0)
}

// Printed when we find a non valid flag. If/else is to match real wget output.
func PrintNotFlag() {
	for _, arg := range NotFlags {
		if strings.HasPrefix(arg, "--") {
			fmt.Printf("wget: unrecognized option '%s'\n", arg)
		} else {
			fmt.Printf("wget: invalid option -- '%s'\n", arg)
		}
	}
	fmt.Println("Usage: go run . [OPTION]... [URL]...\n\nTry `go run  . --help' for more options.")
	os.Exit(0)
}

// Printed if we find a flag but no value (no "=" or string after "=")
func PrintFlagWithoutVal(arg string) {
	fmt.Printf("wget: option '%s' requires an argument\n", arg)
	fmt.Println("Usage: go run . [OPTION]... [URL]...\n\nTry `go run  . --help' for more options.")
	os.Exit(0)
}

// Printed when we find invalid rate number (string).
func PrintInvalidRate(arg string) {
	fmt.Printf("wget: --rate-limit: Invalid byte value ‘%s’\n", arg)
	os.Exit(0)
}

// Print help when -h or --help is found.
func PrintHelp() {
	fmt.Println("Golang wget, a non-interactive network retriever.\nUsage: wget [OPTION]... [URL]...\n\n  -B\t\t\tgo to background after startup\n  -O=FILE\t\twrite documents to FILE\n  -P=PREFIX\t\tsave files to PREFIX/..\n  --rate-limit=RATE\tlimit download rate to RATE\n  -i=FILE\t\tdownload URLs found in FILE\n  --mirror\t\tdownload an entire website\n    -R,  --reject=LIST\tIgnore suffixes in LIST when mirroring.\n    -X,  --exclude=LIST\tIgnore paths in LIST when mirroring\n    --convert-links\tconvert mirrored website links to local links.\n\nSend bug reports, questions, discussions to <https://github.com/kinoz01>")
	os.Exit(0)
}
