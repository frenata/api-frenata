package files

import (
	"html/template"
	"log"
	"net/http"
)

// Not sure why this needs to be /files rather than /files/
const ROUTE string = "/files"

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Method: ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("static/fileform.gtpl")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseMultipartForm(32 << 20)

		_, header, err := r.FormFile("upload")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(getFileMetadata(header)))
	}
}
