package main

import (
	"headers"
	"log"
	"net/http"
	"os"
	"timestamp"
)

func main() {
	// get bound port of host system
	port := os.Getenv("PORT")

	http.HandleFunc(timestamp.ROUTE, timestamp.Handler)
	http.HandleFunc("/headers/", headers.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
