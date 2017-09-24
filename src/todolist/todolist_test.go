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
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respString := string(body)

	require.Equal(t, 404, resp.StatusCode)
	require.Contains(t, strings.ToLower(respString), "not found")
}

func TestRoot(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respString := string(body)

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "Welcome!\n", respString)
}

func TestTodoIndex(t *testing.T) {
	repository = repo.NewInMemoryRepo()
	bytes, _ := json.Marshal(repository.FindAll())
	expected := string(bytes) + "\n"

	req := httptest.NewRequest("GET", "http://example.com/todos", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respString := string(body)

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, expected, respString)
}

func TestTodoShow(t *testing.T) {
	repository = repo.NewInMemoryRepo()
	bytes, _ := json.Marshal(repository.Find("0"))
	expected := string(bytes) + "\n"

	req := httptest.NewRequest("GET", "http://example.com/todos/0", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respString := string(body)

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, expected, respString)
}

func TestTodoCreate(t *testing.T) {
	repository = repo.NewInMemoryRepo()
	request, _ := json.Marshal(spi.Todo{Name: "Test the application"})
	response, _ := json.Marshal(spi.Todo{Name: "Test the application", Id: "3"})

	req := httptest.NewRequest("POST", "http://example.com/todos", bytes.NewReader(request))
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	respString := string(body)

	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, string(response)+"\n", respString)
}
