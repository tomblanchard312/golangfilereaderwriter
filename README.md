# FileReaderWriter  

FileReaderWriter is a simple Go project that demonstrates reading and writing files.  

## Project Structure  

The project is structured as follows:  

filereaderwriter/  
|-- main.go  
|-- filereader/  
| |-- filereader.go  
|-- filewriter/  
| |-- filewriter.go  
|-- README.md  


- `main.go`: The entry point of the program.    
- `filereader/filereader.go`: Package for reading file content.  
- `filewriter/filewriter.go`: Package for writing file content with additional information.  
- `README.md`: This documentation file.  

## Usage  

1. Clone the repository:
   ```
	git clone https://github.com/tomblanchard312/file_readerwritersamplego.git  
	cd filereaderwriter
   ```

### Run the main program:  

	go run main.go
* This will read the content of the example file (example.txt), display it, and write it back with additional user and date information.  

### File Structure  
filereader  
filereader.go  
* This package contains a ReadFile function that reads the contents of a specified file and returns them as a string.  

filewriter  
filewriter.go  
* This package contains a WriteToFile function that writes content to a file and appends the system user's name and the current date.  


## License  
*This project is licensed under the **[MIT License](https://mit-license.org/)**.
