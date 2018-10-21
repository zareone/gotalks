package main

import (
	"fmt"
	"io"
	"net/http"
)

const port = 8080

func main() {
	fmt.Printf("Listening to requests on port %d\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "It works!")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
