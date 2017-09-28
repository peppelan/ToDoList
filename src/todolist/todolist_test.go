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
	repository = repo.NewRepo()

	repository.Create(spi.Todo{Name: "Prepare interview"})
	repository.Create(spi.Todo{Name: "Do not mess up"})

	bytes, _ := json.Marshal(repository.FindAll())
	expected := string(bytes) + "\n"

	code, resp := serveRequest("GET", "http://example.com/todos")

	require.Equal(t, 200, code)
	require.Equal(t, expected, *resp)
}

func TestTodoShow(t *testing.T) {
	repository = repo.NewRepo()

	repository.Create(spi.Todo{Name: "Prepare interview"})

	bytes, _ := json.Marshal(repository.Find("1"))
	expected := string(bytes) + "\n"

	code, resp := serveRequest("GET", "http://example.com/todos/1")

	require.Equal(t, 200, code)
	require.Equal(t, expected, *resp)
}

func TestTodoShowInvalid(t *testing.T) {
	repository = repo.NewRepo()

	code, _ := serveRequest("GET", "http://example.com/todos/gigi")

	require.Equal(t, 404, code)
}

func TestTodoCreate(t *testing.T) {
	repository = repo.NewRepo()
	request, _ := json.Marshal(spi.Todo{Name: "Test the application"})

	req := httptest.NewRequest("POST", "http://example.com/todos", bytes.NewReader(request))
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()
	location := resp.Header.Get("Location")

	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, "/todos/1", location)
}

func TestTodoCreateInvalid(t *testing.T) {
	repository = repo.NewRepo()

	code, _ := serveRequest("POST", "http://example.com/todos")

	require.Equal(t, 422, code)
}

func TestTodoDelete(t *testing.T) {
	repository = repo.NewRepo()
	repository.Create(spi.Todo{Name: "Prepare interview"})
	repository.Create(spi.Todo{Name: "Do not mess up"})

	code, _ := serveRequest("DELETE", "http://example.com/todos/1")
	require.Equal(t, 200, code)
	require.Equal(t, 1, len(repository.FindAll()))
}

func TestTodoDeleteInvalid(t *testing.T) {
	repository = repo.NewRepo()
	repository.Create(spi.Todo{Name: "Prepare interview"})
	repository.Create(spi.Todo{Name: "Do not mess up"})

	code, _ := serveRequest("DELETE", "http://example.com/todos/3")
	require.Equal(t, 404, code)
	require.Equal(t, 2, len(repository.FindAll()))
}

func TestTodoUpdate(t *testing.T) {
	repository = repo.NewRepo()
	repository.Create(spi.Todo{Name: "Prepare interview"})

	request, _ := json.Marshal(spi.Todo{Name: "Test the application"})

	req := httptest.NewRequest("PUT", "http://example.com/todos/1", bytes.NewReader(request))
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestTodoNotFound(t *testing.T) {
	repository = repo.NewRepo()
	request, _ := json.Marshal(spi.Todo{Name: "Test the application"})

	req := httptest.NewRequest("PUT", "http://example.com/todos/1", bytes.NewReader(request))
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()

	require.Equal(t, 404, resp.StatusCode)
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
