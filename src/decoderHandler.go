// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // Handle POST requests for decoding encoded strings
// func decoderHandler(w http.ResponseWriter, r *http.Request) {
// 	// Only allow POST requests
// 	if r.Method != http.MethodPost {
// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Parse form data and handle any parsing errors
// 	if err := r.ParseForm(); err != nil {
// 		http.Error(w, "Malformed form data", http.StatusBadRequest)
// 		return
// 	}

// 	// Get the encoded string from form data
// 	encodedString := r.FormValue("encodedString")

// 	// Attempt to decode the encoded string
// 	decodedString, err := Decoder(encodedString)

// 	// If there's an error, return HTTP 400 with an error message
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)                // Set HTTP status code to 400
// 		fmt.Fprintf(w, "Malformed encoded string: %v", err) // Provide error message
// 		return
// 	}

// 	// If decoding is successful, return HTTP 202 and optionally render a template
// 	w.WriteHeader(http.StatusAccepted)                  // Set HTTP status code to 202
// 	fmt.Fprintf(w, "Decoded result: %s", decodedString) // Output the decoded result
// }

package main

import (
	"net/http"
)

// Handler for decoding operations
func decoderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Use the common function to handle decoding
	processRequest(w, r, "decode", false)
}
