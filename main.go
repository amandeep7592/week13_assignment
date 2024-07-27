package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Fprintf(w, "Error getting hostname: %v", err)
        return
    }

    // HTML content with simple styles
    htmlContent := fmt.Sprintf(`
        <html>
        <head>
            <style>
                body {
                    background-color: #f0f0f0;
                    font-family: Arial, sans-serif;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    margin: 0;
                }
                .container {
                    text-align: center;
                    background-color: #fff;
                    padding: 20px;
                    box-shadow: 0 0 10px rgba(0,0,0,0.1);
                    border-radius: 8px;
                }
                h1 {
                    color: #333;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Hostname: %s</h1>
            </div>
        </body>
        </html>
    `, hostname)

    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, htmlContent)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on port 80")
    if err := http.ListenAndServe(":80", nil); err != nil {
        fmt.Printf("Error starting server: %v\n", err)
    }
}
