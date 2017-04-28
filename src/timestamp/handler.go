package timestamp

import (
	"io"
	"net/http"
	"strings"
)

const HOWTO string = `
Send a requested date or timestamp to /your-request to receive a JSON response back with both the unix epoch and human readable time.

For example:
https://api.frenata.net/timestamp/November 5, 2017

will deliver the response:

{
	"unix": "1509840000",
	"natural": "5 November 2017"
}

Several date formats are supported, try experimenting!

Source code can be found at:
https://github.com/frenata/api-frenata
`

const ROUTE string = "/timestamp/"

// Converts the passed value to a time and writes a JSON response
// including the natural and unix times.
func Handler(w http.ResponseWriter, r *http.Request) {
	// write the instructions
	if r.URL.String() == ROUTE {
		io.WriteString(w, HOWTO)
		return
	}

	// write the JSON
	request := strings.TrimPrefix(r.URL.String(), ROUTE)
	response := GetTimeResponse(request).String()
	io.WriteString(w, response)
}
