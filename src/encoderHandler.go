package main

import (
	"net/http"
)

// Handler for decoding operations
func encoderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Determine if multi-line is selected
	multi := r.FormValue("multi") == "true"

	// Use the common function to handle encoding
	processRequest(w, r, "encode", multi)
}
