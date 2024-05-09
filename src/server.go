package main

import (
	"fmt"
	"net/http"
)

// Handler to serve the main form to the user
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the main HTML form
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<title>Go Encoder/Decoder</title>
			<link rel="stylesheet" type="text/css" href="/static/index.css"> <!-- Linking to external CSS -->
		</head>
		<body>
			<h1>Encoder/Decoder WebApp</h1>
			<form method="post" action="/process">
				<div class="input-group"> <!-- Wrapper for label and textarea -->
					<label for="inputText">Input Text:</label><br> <!-- Label for textarea -->
					<textarea id="inputText" name="input" rows="4" cols="50"></textarea> <!-- Textarea field -->
				</div>

				<div class="form-row"> <!-- Flex container for horizontal alignment -->
					<label for="mode">Mode:</label>
					<select name="mode">
						<option value="decode">Decode</option>
						<option value="encode">Encode</option>
					</select>

					<label for="multi">Multi-line:</label>
					<input type="checkbox" name="multi" value="true">

					<button id="process" type="submit">Process</button>
				</div>
			</form>
		</body>
		</html>
	`)
}
