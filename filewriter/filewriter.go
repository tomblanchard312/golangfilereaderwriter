// filewriter/filewriter.go
package filewriter

import (
	"fmt"
	"os"
	"os/user"
	"time"
)

// WriteToFile writes content to a file and appends the system user's name and the current date.
func WriteToFile(fileName, content string) error {
	// Get the current system user
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	// Create a formatted string with the user's name and current date
	userAndDate := fmt.Sprintf("\n\nWritten by: %s\nDate: %s\n", currentUser.Username, time.Now().Format(time.RFC3339))

	// Append the user's name and date to the content
	content += userAndDate

	// Write the updated content to the file
	err = os.WriteFile(fileName, []byte(content), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
