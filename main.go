package main

import (
	"github.com/frenata/api-frenata/files"
	"github.com/frenata/api-frenata/headers"
        "html/template"
	"github.com/frenata/api-frenata/images"
	"log"
	"github.com/frenata/api-frenata/mytime"
	"net/http"
	"os"
	"github.com/frenata/api-frenata/timestamp"
	"github.com/frenata/api-frenata/tiny"
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
