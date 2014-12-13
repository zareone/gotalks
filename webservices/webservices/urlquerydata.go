package main

import (
	"fmt"
	"net/http"
	"strings"
)

// START OMIT
func myHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Movie search\n\nParameters\n")

	for k, vs := range r.URL.Query() {
		fmt.Fprintf(w, "%s: %s\n", strings.Title(k), strings.Join(vs, ", "))
	}
}

// END OMIT

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", mux)
}
