# art decoder/encoder console app

Welcome to the Art Project repository! This project is a Go-based tool designed to encode and decode text using custom rules. The project includes multiple files for encoding, decoding, and multi-line operations.

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Getting Started](#getting-started)
   - [Installation](#installation)
   - [Usage](#usage)

## Overview
The Art Project is a simple Go program with the following functionalities:

- Encode text using a custom algorithm.
- Decode encoded text to its original form.
- Process multi-line text files for encoding or decoding.

The project is structured in multiple Go files:

- `main.go`: Entry point for the application.
- `encoder.go`: Handles text encoding.
- `decoder.go`: Handles text decoding.
- `multiline.go`: Manages multi-line encoding/decoding operations.

## Features
- Custom text encoding and decoding.
- Multi-line text file processing.
- Command-line interface with multiple options.

## Getting Started

### Installation
To run the project, you need Go installed on your system. If you don't have it installed, [download Go](https://golang.org/dl/) and follow the installation instructions for your operating system.

### Usage
After installing Go, navigate to the project directory and run the following commands based on your use case:

#### Decoding
To decode a single input string:

```bash
go run . "<encoded_string>"
```

#### Encoding
To ecode a single input string:

```bash
go run . -e "<text_to_encode>"
```

#### Multi-line Decoding
To decode an entire file:

```bash
go run . -m <input_file_location>
```

#### Multi-line Encoding
To decode a single input string:

```bash
go run . -m -e <input_file_location>
```



# art encoder/decoder interface
This application allows users to encode and decode text strings. It supports both single-line and multi-line processing for encoding and decoding operations. The application provides a web interface for manual input and a REST API for automated processing.

## Features
- **Encoding**: Compress text strings by encoding repeating characters.
- **Decoding**: Decompress encoded strings back to their original form.
- **Multi-Line Support**: Process multi-line inputs for encoding and decoding.
- **HTTP Endpoints**: Includes an API for encoding and decoding operations.

## Project Structure
- `src/decoder.go`: Decodes encoded text strings.
- `src/encoder.go`: Encodes text strings.
- `src/multiline.go`: Handles multi-line processing for encoding and decoding operations.
- `src/server.go`: Contains HTTP handlers and the web server initialization logic.
- `src/main.go`: Entry point for the application; starts the HTTP server.
- `src/decoderHandler.go`: Defines the HTTP endpoint for decoding encoded strings. It processes POST requests to decode strings and returns appropriate HTTP status codes for success (HTTP 202) or errors (HTTP 400).
- `src/processHandler.go`: Handles form submissions from the web interface for both encoding and decoding. It processes POST requests, applies the desired encoding or decoding operation, and returns the result with the correct HTTP status codes.
- `src/processRequest.go`: Contains a common function for processing encoding and decoding operations. This function is used by both `processHandler` and `decoderHandler`, providing consistent error handling and HTTP status code management.
- `static/`: Contains static files such as CSS and HTML templates for the web interface.

## Getting Started
### Prerequisites
- Go 1.16 or later.
- A web browser to access the application.
- A terminal or command line for API requests.

### Installation
1. Clone the repository to your local environment:
   ```bash
   git clone <repository_url>
   cd <repository_folder>

### To run the application
1. cd <src>
```bash
go run *.go
```
#### To encode or decode a string
1. Type or Paste your strings/texts into the textarea (Input Text box)
2. Then from Mode, select Decode or Encode depending on what action you need to perform
3. Then click process and it will take you to the process page, where it will display the result

#### To encode or decode a multi-line string
1. Type or Paste the location of your text file <file_location>
2. Then from Mode, select Decode or Encode depending on what action you need to perform
3. Then click process and it will take you to the process page, where it will display the result

#### To stop the application
Press ctrl+C from the terminal, and it will terminate the application.