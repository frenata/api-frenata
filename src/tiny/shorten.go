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

const DB string = "redirect.json"

var translate map[string]string

type ShortenerResponse struct {
	Original *string `json:"original_url"`
	Short    *string `json:"short_url"`
}

func NewResponse(original string) ShortenerResponse {
	buff := []byte(original)
	hash := fmt.Sprintf("%x", buff)
	hash = hash[len(hash)-6:]

	short := root + "/tiny/" + hash
	translate[short] = original

	return ShortenerResponse{&original, &short}
}

func (sr ShortenerResponse) String() string {
	json, _ := json.MarshalIndent(sr, "", "    ")
	return string(json)
}

func init() {
	translate = make(map[string]string)

	buff, err := ioutil.ReadFile(DB)
	if err != nil {
		log.Println("redirect db does not exist")
		return
	}

	for _, line := range bytes.Split(buff, []byte("\n")) {
		if len(line) == 0 {
			return
		}
		log.Print(line)
		var sr ShortenerResponse
		err := json.Unmarshal(line, &sr)
		if err != nil {
			panic(err)
		}
		translate[*sr.Short] = *sr.Original
	}
}

func Shorten(requestedURI string) string {
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
