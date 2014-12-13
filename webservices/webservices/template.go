package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// START_TPL OMIT
var html = `
<html>
  <head><title>Our first HTML template</title></head>
  <body>
    <h1>Howdy {{.Name}}!</h1>

    <h3>Current time is {{.Now.Format "15:04:05" }}</h3>
  </body>
</html>`

// END_TPL OMIT

// START_HND OMIT
var tpl = template.Must(template.New("tpl").Parse(html))

func myHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
		Now  time.Time
	}{getParam(r, "name"), time.Now().Local()}

	handleErr(tpl.Execute(w, data))
}

// END_HND OMIT

func getParam(r *http.Request, name string) string {
	ps := r.URL.Query()
	vs := ps[name]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
