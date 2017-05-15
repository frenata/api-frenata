package mytime

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

// possible layouts for parsing date
var layouts = []string{
	"3:04pm utc-07",
	"15:04 utc-07",
	"15 utc-07",
	"3pm utc-07",
	"3:04pm-07",
	"15:04-07",
	"15-07",
	"3pm-07",
	"3pm",
	"15",
}

// TimeResponse represents a JSON response
type TimeResponse struct {
	TheirTime *string `json:"their_time"`
	MyTime    *string `json:"mytime"`
}

// NewTimeResponse generates a JSON response from a given time
func NewTimeResponse(t time.Time, offset int) TimeResponse {
	theirTime := t.Format("3:04pm utc-07")

	zone := time.FixedZone("user", offset*60*60)

	myTime := t.In(zone).Format("3:04pm utc-07")
	return TimeResponse{&theirTime, &myTime}
}

// Pretty prints a TimeResponse in JSON format
func (tr TimeResponse) String() string {
	json, _ := json.MarshalIndent(tr, "", "    ")
	return string(json)
}

// Given a URL request, generates a JSON response string
func getTime(request string) string {
	// process URL to get request
	if strings.HasSuffix(request, "/") {
		request = request[:len(request)-1]
	}
	request = strings.Replace(request, "%20", " ", -1)
	strs := strings.Split(request, "/")
	offset, err := strconv.Atoi(strs[0])
	if err != nil || len(strs) < 2 {
		return TimeResponse{}.String()
	}
	request = strs[1]

	log.Println("request: " + request)
	log.Println("offset: ", offset)

	// get a time from the request
	reqTime, err := parseTime(request)
	if err != nil {
		return TimeResponse{}.String()
	}

	return NewTimeResponse(*reqTime, offset).String()
}

// Given a request, try to parse it into a date
// Returns an error if no way to parse it is found
func parseTime(request string) (*time.Time, error) {
	// add padding
	if len(request) < 2 {
	} else if request[len(request)-2] == '-' || request[len(request)-2] == '+' {
		log.Println("add padding")
		request = request[:len(request)-1] + "0" + request[len(request)-1:]
		log.Println("new request", request)
	}

	// First try the human formats
	for _, layout := range layouts {
		timestamp, err := time.Parse(layout, request)
		if err == nil {
			return &timestamp, nil
		}
	}

	return nil, errors.New("unrecognized time format")
}
