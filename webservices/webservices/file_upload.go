package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// START_TPL OMIT
var html = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>File upload</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap-theme.min.css">
  </head>
  <body>
    <div class="container">
      <h1>File upload</h1>

      <form role="form" action="/" method="post" enctype="multipart/form-data">
        <div class="form-group">
          <label for="fileTitle">Title</label>
          <input type="text" class="form-control" name="title" id="fileTitle" placeholder="Please, enter a title for the file">
        </div>

        <div class="form-group">
          <label for="fileUpload">File</label>
          <input type="file" name="upload" id="fileUpload">
        </div>

        <button type="submit" class="btn btn-default">Upload</button>
      </form>

      {{range $f := .Files}}
        <p>File: {{$f.Title}} <a href="/download/{{$f.ID}}">download</a></p>
      {{else}}
        <p>No files uploaded yet</p>
      {{end}}

    </div>
  </body>
</html>`

// END_TPL OMIT

var tpl = template.Must(template.New("tpl").Parse(html))

type File struct {
	ID          string
	Title       string
	ContentType string
	Content     []byte
}

var files = make(map[string]*File, 0)

// START_HND OMIT
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handleErr(uploadFile(r))
	}

	data := struct {
		Files map[string]*File
	}{files}

	handleErr(tpl.Execute(w, data))
}

// END_HND OMIT

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path[1:], "/")
	if len(urlParts) < 2 {
		panic("Invalid ID in URL")
	}
	ID := urlParts[1]
	f := files[ID]
	if f == nil {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write(f.Content)
	if err != nil {
		panic("Error writing file")
	}
}

func uploadFile(r *http.Request) error {
	f, fi, err := r.FormFile("upload")
	if err != nil {
		return err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	ID, err := newUUID()
	if err != nil {
		return err
	}

	upload := &File{
		ID:          ID,
		Title:       r.FormValue("title"),
		ContentType: fi.Header.Get("Content-Type"),
		Content:     content,
	}

	files[upload.ID] = upload
	return nil
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// newUUID generates a random UUID according to RFC 4122
// copied from http://play.golang.org/p/4FkNSiUDMg
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", uploadHandler)
	mux.HandleFunc("/download/", downloadHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
