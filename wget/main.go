package main

import (
	"os"
	"sync"
	"wget/wget"
)

func main() {

	if len(os.Args) == 1 {
		wget.PrintMissingURL()
	}
	wget.ParseArgs(os.Args[1:])
	if !wget.Mirror {
		wget.Run()
	} else {
		var wg sync.WaitGroup
		wg.Add(1)

		// Parameters
		url := wget.URLs[0]
		rejectList := wget.RejectedSuffixes // Reject image files
		excludeList := wget.ExcludedPaths   // Exclude private directories
		convertLinks := wget.ConvertLinks                   // Convert links to local resources
		logFile, _ := os.Create("mirror.log")
		defer logFile.Close()

		wget.MirrorWebsite(url, rejectList, excludeList, convertLinks, logFile, &wg)

		wg.Wait()
	}
}
