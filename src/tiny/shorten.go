package tiny

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
)

// The flat database file
const DB string = "redirect.json"

// in memory map to translate short URLs to long URLs
var translate map[string]string

// JSON response structure
type ShortenerResponse struct {
	Original *string `json:"original_url"`
	Short    *string `json:"short_url"`
}

// NewResponse creates a short URL from a long URL, adds to the
// translate map, then returns the Response object.
func NewResponse(original string) ShortenerResponse {
	buff := []byte(original)
	hash := fmt.Sprintf("%x", buff)
	hash = hash[len(hash)-6:]

	short := root + "/tiny/" + hash
	translate[short] = original

	return ShortenerResponse{&original, &short}
}

// JSON pretty-print the response
func (sr ShortenerResponse) String() string {
	json, _ := json.MarshalIndent(sr, "", "    ")
	return string(json)
}

// When the library first loads, read from the database and build the translate map from the objects in it.
func init() {
	translate = make(map[string]string)

	buff, err := ioutil.ReadFile(DB)
	if err != nil {
		log.Println("redirect db does not exist")
		return
	}

	log.Println("Reading shortened URL database")
	for _, line := range bytes.Split(buff, []byte("\n")) {
		if len(line) == 0 {
			return
		}
		//log.Print(line)
		var sr ShortenerResponse
		err := json.Unmarshal(line, &sr)
		if err != nil {
			panic(err)
		}
		translate[*sr.Short] = *sr.Original
	}
}

// shorten takes a URL and returns a JSON response containing that URL
// and the mapped short URL.
func shorten(requestedURI string) string {
	//if valid := govalidator.IsHost(requestedURI); !valid {
	if !strings.HasPrefix(requestedURI, "http") {
		requestedURI = "http://" + requestedURI
	}

	u, _ := url.Parse(requestedURI)
	log.Print(u)
	if _, err := net.LookupCNAME(u.Host); err != nil {
		log.Print(err)
		return "{\"error\": \"URL invalid\"}"
	}

	redirect := NewResponse(requestedURI)
	writeDB(redirect)
	return redirect.String()
}

// Appends new responses to the database.
func writeDB(sr ShortenerResponse) {
	f, err := os.OpenFile(DB, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, _ := json.Marshal(sr)
	if _, err = f.Write(append(data, '\n')); err != nil {
		panic(err)
	}
}
