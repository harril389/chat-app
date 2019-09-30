package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("The server is running at 127.0.0.1:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}