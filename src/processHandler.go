package main

import (
	"fmt"
	"net/http"
)

// Process the form submission and return HTTP 202 for success, HTTP 400 for errors
func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Retrieve the form values
	input := r.FormValue("input")
	mode := r.FormValue("mode")
	multi := r.FormValue("multi")

	var result string
	var err error

	if mode == "decode" {
		if multi == "true" {
			result, err = Multiline(input, false) // Process multi-line decode
		} else {
			result, err = Decoder(input) // Process single-line decode
		}
	} else if mode == "encode" {
		if multi == "true" {
			result, err = Multiline(input, true) // Process multi-line encode
		} else {
			result, err = Encoder(input) // Process single-line encode
		}
	}

	// Handle errors, returning HTTP 400 with error message
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Explicitly set HTTP 400
		fmt.Fprintf(w, "Error processing request: %v", err)
		return
	}

	// If the process is successful, return HTTP 202 with the result
	w.WriteHeader(http.StatusAccepted)    // Explicitly set HTTP 202
	fmt.Fprintf(w, "Result:\n%s", result) // Output the result
}
