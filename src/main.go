package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Register HTTP handlers
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) // Serve static files
	http.HandleFunc("/", indexHandler)                                                           // Main form handler

	// Register the /decoder endpoint
	http.HandleFunc("/decoder", decoderHandler)

	// Register the /encoder endpoint
	http.HandleFunc("/encoder", encoderHandler)

	// Start the HTTP server
	fmt.Println("Server is running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
