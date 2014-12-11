package main

import (
	"io"
	"net/http"
)

// START OMIT
func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the default handler")
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the foo handler")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/foo", foo)

	http.ListenAndServe(":8080", mux)
}

// END OMIT
