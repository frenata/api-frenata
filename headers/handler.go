package headers

import "net/http"

const ROUTE string = "/headers/"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(getHeaders(r.Header)))
}
