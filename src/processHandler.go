package main

import (
	"net/http"
)

// Handler for processing form submissions
func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Determine if multi-line is selected
	multi := r.FormValue("multi") == "true"

	// Determine mode (encode or decode)
	mode := r.FormValue("mode")

	// Use the common function to handle processing
	processRequest(w, r, mode, multi)
}
