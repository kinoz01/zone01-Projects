package asciiart

import (
	"fmt"
	"os"
)

func CreateFile(output, outputFile string) {
	// Create a new file or truncate the existing file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
