package main

import (
	"fmt"
	"net/http"
	"os"
)


func mixColors(color1, color2 string) string {
	mixedColor := ""
	switch color1 {
	case "red":
		if color2 == "blue" {
			mixedColor = "purple"
		} else if color2 == "yellow" {
			mixedColor = "orange"
		} else {
			mixedColor = "unknown"
		}
	case "blue":
		if color2 == "red" {
			mixedColor = "purple"
		} else if color2 == "yellow" {
			mixedColor = "green"
		} else {
			mixedColor = "unknown"
		}
	case "yellow":
		if color2 == "red" {
			mixedColor = "orange"
		} else if color2 == "blue" {
			mixedColor = "green"
		} else {
			mixedColor = "unknown"
		}
	default:
		mixedColor = "unknown"
	}
	return mixedColor
}


func colorHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}

	var color1, color2 string
	switch hostname {
	case "server1":
		color1, color2 = "red", "blue"
	case "server2":
		color1, color2 = "blue", "yellow"
	case "server3":
		color1, color2 = "yellow", "red"
	default:
		color1, color2 = "unknown", "unknown"
	}

	mixedColor := mixColors(color1, color2)
	fmt.Fprintf(w, "Hostname: %s\nMixing %s and %s results in: %s", hostname, color1, color2, mixedColor)
}

func main() {
	http.HandleFunc("/", colorHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
