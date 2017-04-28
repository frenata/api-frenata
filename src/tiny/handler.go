package tiny

import (
	"io"
	"log"
	"net/http"
	"strings"
)

const ROUTE string = "/tiny/"

var root string

func Handler(w http.ResponseWriter, r *http.Request) {
	root = r.Host
	if !strings.HasPrefix(root, "localhost") {
		root = "https://" + root
	}
	requestedURI := strings.TrimPrefix(r.URL.String(), ROUTE)
	requestedURI = strings.Replace(requestedURI, ":/", "://", 1)
	log.Println("request: ", requestedURI)

	from := r.Host + r.URL.Path
	if redirect, ok := translate[from]; ok {
		log.Println("redirecting from: ", from)
		log.Println("redirecting to: ", redirect)
		http.Redirect(w, r, redirect, 301)
		return
	}

	response := Shorten(requestedURI)
	io.WriteString(w, response)
}
