package asciiart

import (
	"fmt"
	"os"
)
// takes the content and the file name, create the file and write the content (ascii art) in it.
func CreateFile(output, outputFileName string) {
	// Create a new file or truncate the existing file
	file, err := os.Create(outputFileName)
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
