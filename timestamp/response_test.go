package timestamp

import "testing"

func TestLayouts(t *testing.T) {
	for _, layout := range layouts {
		if _, err := parseTime(layout); err != nil {
			t.Fatalf("%s layout fails to pare itself", layout)
		}
	}
}

func TestEpoch1(t *testing.T) {
	unix := "1"
	natural := "1 January 1970"
	expected := TimeResponse{&unix, &natural}.String()

	if res := getTime("1"); res != expected {
		t.Log(res)
		t.Log(expected)
		t.Fatal("Epoch 1 does not create proper response")
	}
}

func TestBadRequest(t *testing.T) {
	request := "foobar/"

	response := getTime(request)
	blank := TimeResponse{}

	if response != blank.String() {
		t.Fatalf("%s did not generate a null response", request)
	}
}
