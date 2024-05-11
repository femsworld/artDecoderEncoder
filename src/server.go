package main

import (
	"net/http"
)

// Handler to serve the main form to the user
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the main HTML form
	http.ServeFile(w, r, "static/index.html")
}
