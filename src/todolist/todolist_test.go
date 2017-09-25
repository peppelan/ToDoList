package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"todolist/repo"
	"todolist/spi"
)

func TestNonHandledEndpoint(t *testing.T) {
	code, resp := serveRequest("GET", "http://example.com/foo")

	require.Equal(t, 404, code)
	require.Contains(t, strings.ToLower(*resp), "not found")
}

func TestRoot(t *testing.T) {
	code, resp := serveRequest("GET", "http://example.com/")

	require.Equal(t, 200, code)
	require.Equal(t, "Welcome!\n", *resp)
}

func TestTodoIndex(t *testing.T) {
	repository = repo.NewInMemoryRepo()
	bytes, _ := json.Marshal(repository.FindAll())
	expected := string(bytes) + "\n"

	code, resp := serveRequest("GET", "http://example.com/todos")

	require.Equal(t, 200, code)
	require.Equal(t, expected, *resp)
}

func TestTodoShow(t *testing.T) {
	repository = repo.NewInMemoryRepo()
	bytes, _ := json.Marshal(repository.Find("0"))
	expected := string(bytes) + "\n"

	code, resp := serveRequest("GET", "http://example.com/todos/0")

	require.Equal(t, 200, code)
	require.Equal(t, expected, *resp)
}

func TestTodoCreate(t *testing.T) {
	repository = repo.NewInMemoryRepo()
	request, _ := json.Marshal(spi.Todo{Name: "Test the application"})
	expectedResponse, _ := json.Marshal(spi.Todo{Name: "Test the application", Id: "3"})

	req := httptest.NewRequest("POST", "http://example.com/todos", bytes.NewReader(request))
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respString := string(body)
	location := resp.Header.Get("Location")

	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, string(expectedResponse)+"\n", respString)
	require.Equal(t, "/todos/3", location)
}

func TestTodoCreateInvalid(t *testing.T) {
	repository = repo.NewInMemoryRepo()

	code, _ := serveRequest("POST", "http://example.com/todos")

	require.Equal(t, 422, code)
}

func TestTodoDelete(t *testing.T) {
	repository = repo.NewInMemoryRepo()

	code, _ := serveRequest("DELETE", "http://example.com/todos/1")
	require.Equal(t, 200, code)
	require.Equal(t, 1, len(repository.FindAll()))
}

func TestTodoDeleteInvalid(t *testing.T) {
	repository = repo.NewInMemoryRepo()

	code, _ := serveRequest("DELETE", "http://example.com/todos/3")
	require.Equal(t, 404, code)
	require.Equal(t, 2, len(repository.FindAll()))
}

// Serves a given request, returns status code and response body
func serveRequest(method string, target string) (int, *string) {
	req := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)
	res := w.Result()
	resBody, _ := ioutil.ReadAll(res.Body)
	respString := string(resBody)
	return res.StatusCode, &respString
}