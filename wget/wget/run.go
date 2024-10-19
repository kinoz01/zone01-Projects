package wget

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Download URLs strings concurrently while logging the process to either "wget-log" or the stdout.
func Run() {
	var wg sync.WaitGroup

	if Silent {
		logFileName := GetLogFileName()
		logFile, err := os.Create(logFileName)
		if err != nil {
			fmt.Printf("Error creating log file '%s': %v\n", logFileName, err)
			os.Exit(1)
		}
		defer logFile.Close()
		LogOutput = logFile
		fmt.Printf("Output will be written to \"%s\".\n", logFileName)
	} else {
		LogOutput = os.Stdout
	}

	startTime := time.Now()
	fmt.Fprintf(LogOutput, "start at %s\n", startTime.Format("2006-01-02 15:04:05"))

	for _, rawURL := range URLs {
		wg.Add(1)
		go func(rawURL string) {
			defer wg.Done()
			err := DownloadFile(rawURL)
			if err == nil { // Error already logged in the function.
				fmt.Fprintf(LogOutput, "\nDownloaded [%s]\n", rawURL)
			}
		}(rawURL)
	}

	wg.Wait()

	endTime := time.Now()
	fmt.Fprintf(LogOutput, "finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}

// Check if wget-log name is already used and use "wget-log.1" and so on.
func GetLogFileName() string {
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