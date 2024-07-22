package main

import (
	"filereaderwriter/internal/reader"
	"filereaderwriter/internal/writer"
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := "example.txt"

	// Get the current date and time
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("hello\ngo\n%s\n", currentTime)

	// Write to the file (or append if it exists)
	err := writer.WriteOrAppendToFile(filePath, message)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
	fmt.Println("File written/appended successfully")

	// Read from the file
	content, err := reader.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	fmt.Println("File Content:")
	fmt.Println(content)
}
