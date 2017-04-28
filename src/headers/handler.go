package headers

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, GetHeaders(r.Header))
}
