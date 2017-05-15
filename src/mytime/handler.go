package mytime

import (
	"io"
	"net/http"
	"strings"
)

const HOWTO string = `
Request a time to be translated to your TZ like so:

/11:30am utc+04

minutes, am/pm, 'utc' are all optional
`

const ROUTE string = "/mytime/"

func Handler(w http.ResponseWriter, r *http.Request) {
	// write the instructions
	if r.URL.String() == ROUTE {
		io.WriteString(w, HOWTO)
		return
	}

	// write the JSON
	request := strings.TrimPrefix(r.URL.String(), ROUTE)
	response := getTime(request)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
