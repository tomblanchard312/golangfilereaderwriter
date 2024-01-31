// filereader.go
package filereader

import (
	"os"
)

// ReadFile reads the contents of the specified file and returns them as a string.
func ReadFile(fileName string) (string, error) {
	// Use os.ReadFile to read the entire content of the file into a byte slice
	content, err := os.ReadFile(fileName)

	// Check for errors during file reading
	if err != nil {
		// If an error occurs, return an empty string and the error
		return "", err
	}

	// Convert the byte slice to a string and return it along with no errors
	return string(content), nil
}
