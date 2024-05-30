package main

import (
	"fmt"
	"net/http"
)

// Exported function to handle requests
func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for key, values := range query {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", key, value)
		}
	}
}
