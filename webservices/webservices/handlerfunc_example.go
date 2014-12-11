package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// START OMIT
func myHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi from the Go server!")
}

func main() {
	// Let's create a fake request
	r, err := http.NewRequest("GET", "http://localhost:8080", nil)
	handleErr(err)

	// Also create a fake ResponseWriter
	w := httptest.NewRecorder()

	// We do type conversion from our handler function to http.HandlerFunc type
	// which implements the ServeHTTP method defined in http.Handler interface
	adaptedHandler := http.HandlerFunc(myHandler) // HL
	adaptedHandler.ServeHTTP(w, r)

	fmt.Printf("Response: \ncode: %d\nbody: %s", w.Code, w.Body.String())
}

// END OMIT

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
