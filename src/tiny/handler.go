package tiny

import (
	"io"
	"log"
	"net/http"
	"strings"
)

const ROUTE string = "/tiny/"

// for handling localhost vs deployed
var root string

func Handler(w http.ResponseWriter, r *http.Request) {
	root = r.Host
	if !strings.HasPrefix(root, "localhost") {
		root = "https://" + root
	}
	// fix up the requested URL and log it
	requestedURI := strings.TrimPrefix(r.URL.String(), ROUTE)
	requestedURI = strings.Replace(requestedURI, ":/", "://", 1)
	log.Println("request: ", requestedURI)

	// check the translate map to see if this is a previously designated
	// short URL. If so, redirect to the long URL.
	from := root + r.URL.Path
	if redirect, ok := translate[from]; ok {
		log.Println("redirecting from: ", from)
		log.Println("redirecting to: ", redirect)
		http.Redirect(w, r, redirect, 301)
		return
	}

	// get a response from Shorten
	var response string
	if requestedURI == "" {
		response = "Simply add any URL after /tiny/"
	} else {
		response = Shorten(requestedURI)
	}

	// write the response to the browser
	io.WriteString(w, response)
}
