package main

import (
	"fmt"
	"net/http"
)

func decoderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("input")
	multi := r.FormValue("multi") == "true"

	var result string
	var err error
	if multi {
		result, err = Multiline(input, false) // Multi-line decode
	} else {
		result, err = Decoder(input) // Single-line decode
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %v", err), http.StatusBadRequest)
		return
	}

	renderResult(w, "Decode Result", result)
}
