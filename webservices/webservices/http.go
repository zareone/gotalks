package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hi from the Go server!")
	})

	http.ListenAndServe(":8080", nil)
}
