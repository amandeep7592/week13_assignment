package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MixRequest struct {
	Color1 string `json:"color1"`
	Color2 string `json:"color2"`
}

type MixResponse struct {
	ResultColor string `json:"result_color"`
	Message     string `json:"message"`
}

func mixColorsHandler(w http.ResponseWriter, r *http.Request) {
	var mixRequest MixRequest
	err := json.NewDecoder(r.Body).Decode(&mixRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	resultColor, message := mixColors(mixRequest.Color1, mixRequest.Color2)

	response := MixResponse{
		ResultColor: resultColor,
		Message:     message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func mixColors(color1, color2 string) (string, string) {
	colorMixes := map[string]string{
		"red+blue":   "purple",
		"red+yellow": "orange",
		"blue+yellow":"green",
		"blue+red":   "purple",
		"yellow+red": "orange",
		"yellow+blue":"green",
	}

	if color1 == color2 {
		return color1, "Same colors"
	}

	result, exists := colorMixes[color1+"+"+color2]
	if !exists {
		result, exists = colorMixes[color2+"+"+color1]
	}

	if !exists {
		return "unknown", "Unknown color combination"
	}

	return result, "Mixed successfully"
}

func main() {
	http.HandleFunc("/mix", mixColorsHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
