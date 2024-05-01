# art

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