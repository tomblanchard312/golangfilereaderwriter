package writer

import (
	"bufio"
	"os"
)

// WriteOrAppendToFile writes to the file if it doesn't exist, or appends to it if it does.
func WriteOrAppendToFile(filePath string, data string) error {
	var file *os.File
	var err error

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}
	return writer.Flush()
}
