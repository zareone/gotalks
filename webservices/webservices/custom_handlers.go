package main

import (
	"fmt"
	"http"
	"log"
)

type Context struct{}

func (c *Context) Close() {}

func handleError(w http.ResponseWriter, r *http.Request, err error) {}

func panicHandler(w http.ResponseWriter, r *http.Request) {}

func newContext(r *http.Request) *Context { return &Context{} }

// START_CUSTOM_HANDLER OMIT
type MyHandlerFunc func(http.ResponseWriter, *http.Request, *Context) error // HL

func HandlerAdapter(f HandlerFunc) httptreemux.HandlerFunc {
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

// START_MAIN OMIT
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkout", HandlerAdapter(checkoutHandler))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func checkoutHandler(w http.ResponseWriter, r *http.Request, ctx *Context) error {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	ret := map[string]float64{"total": total}
	return writeJSON(w, 200, ret) // HL
}

//END_MAIN OMIT
