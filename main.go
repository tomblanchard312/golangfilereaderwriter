// main.go
package main

import (
	"file_reader/filereader"
	"file_reader/filewriter"
	"fmt"
)

func main() {
	fileName := "example.txt"

	// Read the file
	content, err := filereader.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:")
	fmt.Println(content)

	// Write the file with user and date information
	err = filewriter.WriteToFile(fileName, content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File successfully written with user and date information.")
}
