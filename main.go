package main

import (
	"fmt"
	"net/http"
	"os"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Hostname: %s", hostname)
}

func main() {
	http.HandleFunc("/", hostnameHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
