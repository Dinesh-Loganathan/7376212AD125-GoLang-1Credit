package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Dinesh L")
	})
	http.HandleFunc("/regno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "212AD125")
	})
	http.HandleFunc("/dept", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "AI&DS")
	})
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Black")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
