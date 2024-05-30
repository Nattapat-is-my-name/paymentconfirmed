package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test!")
}

func init() {
	fmt.Println("Initializing...")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", handler2)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
