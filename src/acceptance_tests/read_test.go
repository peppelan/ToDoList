package acceptance_tests

import (
	"testing"
	"flag"
	"net/http"
	"io/ioutil"
	"github.com/stretchr/testify/require"
)

var (
	url = flag.String("url", "http://localhost:8080", "URL to test")
)

func TestRead(t *testing.T) {

	expected := string("Hello, \"/\"")

	resp, err := http.Get(*url)

	defer resp.Body.Close()

	require.Nil(t, err, "Received unexpected error")
	require.Equal(t, 200, resp.StatusCode, "Received unexpected status code")

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	require.Nil(t, err, "Received unexpected error")
	require.Equal(t, expected, bodyString, "Received unexpected response body")
}
