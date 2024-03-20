package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function to serve the webpage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Define the HTML content
		html := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Simple Web Page</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    margin: 0;
                    padding: 0;
                    background-color: #f0f0f0;
                }
                .container {
                    max-width: 800px;
                    margin: 50px auto;
                    padding: 20px;
                    background-color: #fff;
                    border-radius: 5px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }
                h1 {
                    color: #333;
                    border-bottom: 2px solid #333;
                    padding-bottom: 10px;
                }
                p {
                    color: #666;
                }
                .info {
                    background-color: #ffc107;
                    color: #333;
                    padding: 10px;
                    border-radius: 5px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Welcome to My Simple Web Page!</h1>
                <p>This is a basic example of serving a webpage using Go.</p>
                <div class="info">
                    <p>This page is created using Go programming language.</p>
                </div>
				<div class="container">
                <h1>Current Temperature</h1>
                <p>%s</p>
            </div>
            </div>
        </body>
        </html>
        `
		// Write the HTML content to the response
		fmt.Fprintf(w, html)
	})

	// Start the web server on port 8080
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
