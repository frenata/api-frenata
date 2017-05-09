package files

import (
	"html/template"
	"io"
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
		r.ParseForm()

		io.WriteString(w, r.Form["test"][0])
	}
}
