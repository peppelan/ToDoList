package main

import(
	"strings"
	"testing"
	"net/http/httptest"
	"io/ioutil"

	"github.com/stretchr/testify/require"
)

func TestNonHandledEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	Handler().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respString := string(body)

	require.Equal(t, 404, resp.StatusCode)
	require.Contains(t, strings.ToLower(respString), "not found")
}

func TestRoot(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/", nil)
	w := httptest.NewRecorder()
	Handler().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respString := string(body)

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, respString, "Hello, \"/\"")
}