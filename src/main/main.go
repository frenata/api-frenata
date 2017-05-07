package main

import (
	"headers"
	"images"
	"io"
	"log"
	"net/http"
	"os"
	"timestamp"
	"tiny"
)

const WELCOME string = `
Welcome to my collection of API microservices:

	/timestamp/
	/headers/
	/tiny/
	/images/
	...
`

func main() {
	// get bound port of host system
	port := os.Getenv("PORT")

	http.HandleFunc("/", index)
	http.HandleFunc(timestamp.ROUTE, timestamp.Handler)
	http.HandleFunc(headers.ROUTE, headers.Handler)
	http.HandleFunc(tiny.ROUTE, tiny.Handler)
	http.HandleFunc(images.ROUTE, images.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, WELCOME)
}
