package acceptance_tests

import (
	"testing"
	"flag"
	"net/http"
	"io/ioutil"
)

var (
	url = flag.String("url", "http://localhost:8080", "URL to test")
)

func TestRead(t *testing.T) {

	expected := string("Hello, \"/\"")

	resp, err := http.Get(*url)

	if err != nil {
		t.Errorf("Received unexpected error: %s", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Received unexpected status code: %d", resp.StatusCode)
		return
	}

	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err2 != nil {
		t.Errorf("Received unexpected error: %s", err2)
		return
	}

	if bodyString != expected {
		t.Errorf("Received unexpected response: '%s', expected: '%s'", bodyString, expected)
		return
	}
}
