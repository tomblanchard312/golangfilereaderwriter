# FileReaderWriter

[![License](https://img.shields.io/github.com/tomblanchard312/golangfilereaderwriter)](https://github.com/tomblanchard312/golangfilereaderwriter/blob/main/LICENSE)

FileReaderWriter is a simple Go project that demonstrates reading from and writing to files. This project includes:

1. Checking if a file (`example.txt`) exists.
2. Creating and writing to the file if it does not exist.
3. Appending to the file if it already exists.
4. Reading and displaying the file content.
5. Adding timestamps to each write operation.

## Project Structure

The project is structured as follows:

```
filereaderwriter/
├── cmd
│   └── app
│       └── main.go  # Presentation Layer
├── internal
│   ├── reader
│   │   ├── filereader.go  # Business Logic Layer
│   │   └── filereader_test.go
│   ├── writer
│   │   ├── filewriter.go  # Business Logic Layer
│   │   └── filewriter_test.go
│   └── utils
│       ├── error.go  # Utility Layer
│       └── error_test.go
└── README.md
```

- `cmd/app/main.go`: The entry point of the program.
- `internal/reader/filereader.go`: Package for reading file content.
- `internal/writer/filewriter.go`: Package for writing or appending file content with additional information.
- `internal/utils/error.go`: Utility functions for error handling.
- `README.md`: This documentation file.

## Design Patterns

The design pattern used in this project can be considered a combination of the Layered Architecture and Dependency Injection.

Here’s a detailed explanation of the patterns and how they are applied in this project:

Layered Architecture
The project is structured into layers, each responsible for a specific part of the application's functionality:

Presentation Layer: Contains the main.go file which serves as the entry point of the application.
Business Logic Layer: Contains the core logic for reading and writing files.
internal/reader: Responsible for reading file contents.
internal/writer: Responsible for writing or appending to files.
Utility Layer: Contains common utilities that can be used across different layers.
internal/utils: Contains utility functions for error handling.
Dependency Injection
While this project is simple and doesn’t fully demonstrate dependency injection in its classical form, the concept is applied through the use of functions and interfaces to manage dependencies. The main.go file depends on the reader and writer packages but interacts with them through well-defined functions.

### Layered Architecture

The project uses a layered architecture to separate concerns:

- **Presentation Layer**: Handles user interaction and controls the application flow (`cmd/app/main.go`).
- **Business Logic Layer**: Contains the core logic for reading and writing files (`internal/reader` and `internal/writer`).
- **Utility Layer**: Provides utility functions that support the business logic layer (`internal/utils`).

### Dependency Injection

The project uses functions and interfaces to inject dependencies, allowing for loose coupling and easier testing. The `main.go` file interacts with the `reader` and `writer` packages through well-defined functions.

## Usage

1. Clone the repository:

   ```sh
   git clone https://github.com/tomblanchard312/file_readerwritersamplego.git
   cd filereaderwriter

  ```

### Run the main program  

 ```sh
     go build -o gofilereaderwriter.exe ./cmd/app
     go run gofilereaderwriter.go
 ```

This will write to example.txt if it does not exist or append to it if it does. It will then read the content of the file and display it.

### Functionality

 Writing to the File
    -The WriteOrAppendToFile function in filewriter.go:
 --Checks if the file exists.
 --Creates the file and writes to it if it does not exist.
 --Opens the file in append mode and writes to it if it already exists.
 --Adds the current timestamp to each write operation.
 --Reading from the File

 -The ReadFile function in filereader.go:
 --Opens the specified file.
 --Reads its contents and returns them as a string.
 --Error Handling

-The CheckError function in utils/error.go:
--Checks for errors and panics if an error is found.

-Testing
--Testing is included for both the reader and writer functionalities.

-Running Tests
--To run the tests, navigate to the root of your project and execute:

```sh
   go test ./...
```

## License  

*This project is licensed under the **[MIT License](https://github.com/tomblanchard312/GoLangFileReaderWriterSample/edit/main/LICENSE)**.
