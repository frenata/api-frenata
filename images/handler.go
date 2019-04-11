package images

import (
	"io"
	"log"
	"net/http"
	"strings"
)

const HOWTO = `
To search for images, send a request to /images/your-search-here

You can get more results for the same search by adding '?offset=2' to your search

You can also see what search request have been made at /images/latest/
`

const ROUTE string = "/images/"

// Converts the passed value to a time and writes a JSON response
// including the natural and unix times.
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: ", r.URL.String())
	// write the instructions
	if r.URL.String() == ROUTE {
		io.WriteString(w, HOWTO)
		return
	} else if r.URL.String() == ROUTE+"latest/" ||
		r.URL.String() == ROUTE+"latest" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(getSearchHistory()))
		return
	}

	// write the JSON
	request := strings.TrimPrefix(r.URL.String(), ROUTE)
	response := imageSearch(request)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
