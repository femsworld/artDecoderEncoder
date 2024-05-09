package main

import (
	"fmt"
	"net/http"
)

// processRequest handles decoding or encoding operations and returns appropriate HTTP status codes.
func processRequest(w http.ResponseWriter, r *http.Request, mode string, isMultiLine bool) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Get the input data to process
	input := r.FormValue("input")

	var result string
	var err error

	// Select operation based on mode and multi-line flag
	if mode == "decode" {
		if isMultiLine {
			result, err = Multiline(input, false) // Multi-line decode
		} else {
			result, err = Decoder(input) // Single-line decode
		}
	} else if mode == "encode" {
		if isMultiLine {
			result, err = Multiline(input, true) // Multi-line encode
		} else {
			result, err = Encoder(input) // Single-line encode
		}
	}

	// Handle errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Set HTTP 400 for errors
		fmt.Fprintf(w, "Error processing request: %v", err)
		return
	}

	// Successful processing
	w.WriteHeader(http.StatusAccepted)    // Set HTTP 202 for success
	fmt.Fprintf(w, "Result:\n%s", result) // Output the result
}
