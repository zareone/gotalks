package main

import (
	"fmt"
	"log"
	"net/http"
)

type Context struct{}

func (c *Context) Close() {}

func handleError(w http.ResponseWriter, r *http.Request, err error) {}

func panicHandler(w http.ResponseWriter, r *http.Request) {}

func newContext(r *http.Request) *Context { return &Context{} }

// START_CUSTOM_HANDLER OMIT
type MyHandlerFunc func(http.ResponseWriter, *http.Request, *Context) error // HL

func WrapHandler(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer panicHandler(w, r) // Catch any panics that could happen in the hanlder

		ctx, err := newContext(r)
		defer ctx.close()

		if err != nil {
			panic(fmt.Sprintf("error creating a new context: %v", err))
		}

		// Run the wrapped handler
		err = f(w, r, ctx) // HL
		if err != nil {
			handleError(w, r, err)
			return
		}
	}
}

// END_CUSTOM_HANDLER OMIT

// START_HTTPERR OMIT
type HTTPError struct {
	Code    int
	Message string
}

func NewHTTPError(code int, message string) *HTTPError {
	if message == "" {
		message = http.StatusText(code)
	}
	return &HTTPError{Code: code, Message: message}
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

// END_HTTPERR OMIT

// START_MAIN OMIT

var ErrMethodNotAllowed = NewHTTPError(405, "") // HL

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkout", WrapHandler(checkoutHandler)) // HL
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func checkoutHandler(w http.ResponseWriter, r *http.Request, ctx *Context) error {
	if r.Method != "POST" {
		return ErrMethodNotAllowed // HL
	}
	ret := map[string]float64{"total": 545.0}
	return writeJSON(w, 200, ret) // HL
}

//END_MAIN OMIT
