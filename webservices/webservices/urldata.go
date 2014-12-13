package main

import (
	"fmt"
	"net/http"
	"strings"
)

// START OMIT
func myHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Request URL path: %s\n", r.URL.Path)

	for i, part := range strings.Split(r.URL.Path[1:], "/") {
		fmt.Fprintf(w, "URL part #%d: %s\n", i, part)
	}
}

// END OMIT

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", mux)
}
