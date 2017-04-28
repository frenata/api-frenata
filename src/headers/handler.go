package headers

import (
	"io"
	"net/http"
)

const ROUTE string = "/headers/"

func Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, GetHeaders(r.Header))
}
