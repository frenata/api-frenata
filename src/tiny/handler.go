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

	from := root + r.URL.Path
	if redirect, ok := translate[from]; ok {
		log.Println("redirecting from: ", from)
		log.Println("redirecting to: ", redirect)
		http.Redirect(w, r, redirect, 301)
		return
	}

	var response string
	if requestedURI == "" {
		response = "Simply add any URL after /tiny/"
	} else {
		response = Shorten(requestedURI)
	}

	io.WriteString(w, response)
}
