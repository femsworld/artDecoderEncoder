package main

import (
	"fmt"
	"net/http"
)

// Handler for decoding operations
func encoderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("input")
	multi := r.FormValue("multi") == "true"

	var result string
	var err error
	if multi {
		result, err = Multiline(input, true) // Multi-line encode
	} else {
		result, err = Encoder(input) // Single-line encode
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %v", err), http.StatusBadRequest)
		return
	}

	renderResult(w, "Encode Result", result)
}
