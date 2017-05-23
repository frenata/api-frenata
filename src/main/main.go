package main

import (
	"files"
	"headers"
	"html/template"
	"images"
	"log"
	"mytime"
	"net/http"
	"os"
	"timestamp"
	"tiny"
)

func main() {
	// get bound port of host system
	port := os.Getenv("PORT")

	http.HandleFunc("/", index)
	http.HandleFunc(timestamp.ROUTE, timestamp.Handler)
	http.HandleFunc(mytime.ROUTE, mytime.Handler)
	http.HandleFunc(headers.ROUTE, headers.Handler)
	http.HandleFunc(tiny.ROUTE, tiny.Handler)
	http.HandleFunc(images.ROUTE, images.Handler)
	http.HandleFunc(files.ROUTE, files.Handler)
	http.HandleFunc("/files/", files.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
