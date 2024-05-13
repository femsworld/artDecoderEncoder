package main

import (
	"html/template"
	"net/http"
)

// Handler to serve the main form to the user
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the main HTML form
	http.ServeFile(w, r, "static/index.html")
}

func renderResult(w http.ResponseWriter, title, result string) {
	tmpl, err := template.New("result").Parse(resultTemplate)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	data := map[string]string{
		"Title":  title,
		"Result": result,
	}
	w.WriteHeader(http.StatusAccepted)
	tmpl.Execute(w, data)
}

const resultTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #e6e6fa; /* Light purple background color */
        }
        h1 {
            color: #333;
        }
        .result-textarea {
            width: 75%;
            height: 15rem;
            border-radius: 0.5rem;
            border: 1px solid #ccc;
            padding: 10px;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <h1>{{.Title}}</h1>
    <textarea id="resultText" class="result-textarea" readonly>{{.Result}}</textarea>
</body>
</html>
`
