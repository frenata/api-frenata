package images

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var latest SearchHistory = make(SearchHistory, 0)

// ImageData represents JSON input from Flickr
type ImageData struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
	Server string `json:"server"`
	Title  string `json:"title"`
	Farm   int    `json:"farm"`
}

func (ir ImageData) String() string {
	json, _ := json.MarshalIndent(ir, "", "    ")
	return string(json)
}

type ImageDataArray []ImageData

func (ira ImageDataArray) String() string {
	json, _ := json.MarshalIndent(ira, "", "    ")
	return string(json)
}

// ImageResponse represents JSON output
type ImageResponse struct {
	Uri       string `json:"uri"`
	Thumbnail string `json:"thumbnail"`
	Text      string `json:"text"`
}

func (ir ImageResponse) String() string {
	json, _ := json.MarshalIndent(ir, "", "    ")
	return string(json)
}

func NewImageResponse(data ImageData) ImageResponse {
	ir := ImageResponse{}

	uri := fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s",
		data.Farm,
		data.Server,
		data.Id,
		data.Secret)
	ir.Uri = uri + ".jpg"
	ir.Thumbnail = uri + "_s.jpg"
	ir.Text = data.Title

	return ir
}

type ImageResponseArray []ImageResponse

func (ira ImageResponseArray) String() string {
	json, _ := json.MarshalIndent(ira, "", "    ")
	return string(json)
}

// Search represents a specific search request
type Search struct {
	Term string `json:"term"`
	When string `json:"when"`
}

func NewSearch(term string) Search {
	return Search{term, time.Now().String()}
}

func (s Search) String() string {
	json, _ := json.MarshalIndent(s, "", "    ")
	return string(json)
}

type SearchHistory []Search

func (sh SearchHistory) String() string {
	json, _ := json.MarshalIndent(sh, "", "    ")
	return string(json)
}

func imageSearch(request string) string {
	params := strings.Split(request, "?")
	text := params[0]
	log.Println("Searching Flickr for: ", text)
	addSearch(text)

	page := 1
	if len(params) > 1 && strings.HasPrefix(params[1], "offset=") {
		n, err := strconv.Atoi(strings.TrimPrefix(params[1], "offset="))
		if err == nil {
			page = n
			log.Println("Offsetting search to page ", page)
		}
	}

	apikey := "7c16845b11cf78f1c92fc825345df211"
	url := fmt.Sprintf(
		"https://api.flickr.com/services/rest/?method=flickr.photos.search&api_key=%s&text=%s&page=%d&format=json&nojsoncallback=1",
		apikey,
		text,
		page)

	client := http.Client{Timeout: time.Second * 20}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var objmap map[string]*json.RawMessage
	json.Unmarshal(body, &objmap)
	var photomap map[string]*json.RawMessage
	json.Unmarshal(*objmap["photos"], &photomap)
	ida := make(ImageDataArray, 0)
	json.Unmarshal(*photomap["photo"], &ida)

	ira := make(ImageResponseArray, 0)
	for _, v := range ida {
		ira = append(ira, NewImageResponse(v))
	}

	return ira.String()
}

func getSearchHistory() string {
	return latest.String()
}

func addSearch(text string) {
	latest = append([]Search{NewSearch(text)}, latest...)
	if len(latest) > 10 {
		latest = latest[:10]
	}
}
