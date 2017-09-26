package acceptance_tests

import (
	"encoding/json"
	"flag"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"todolist/spi"
)

var (
	url = flag.String("url", "http://localhost:8080", "URL to test")
)

func TestRead(t *testing.T) {

	expected := string("Welcome!\n")
	testRequest(t, "GET", "", "", 200, &expected)
}

func TestFullCrudSession(t *testing.T) {
	myTodo := spi.Todo{Name: "Test the application"}
	request, _ := json.Marshal(myTodo)

	// Create
	myTodoAddress := testRequest(t, "POST", "/todos", string(request), 201, nil)
	println("Working with " + myTodoAddress)
	myTodo.Id = strings.Replace(myTodoAddress, "/todos/", "", -1)

	// Read
	checkTodo(t, myTodo)

	// Update
	testRequest(t, "POST", "/todos", string(request), 201, nil)
	updatedTodo := spi.Todo{Id: myTodo.Id, Name: "Test the application a bit more"}
	myTodoJson, _ := json.Marshal(updatedTodo)
	testRequest(t, "PUT", myTodoAddress, string(myTodoJson), 200, nil)
	checkTodo(t, updatedTodo)

	// Delete
	testRequest(t, "DELETE", myTodoAddress, "", 200, nil)
	testRequest(t, "GET", myTodoAddress, "", 404, nil)
}

// Serves a given request, returns status code and response body
func testRequest(t *testing.T,
	method string,
	target string,
	reqBody string,
	expectedRetCode int,
	expectedResBody *string) string {

	request, _ := http.NewRequest(method, *url+target, strings.NewReader(reqBody))
	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()

	require.Nil(t, err, "Received unexpected error")
	require.Equal(t, expectedRetCode, resp.StatusCode, "Received unexpected status code")

	resBody, _ := ioutil.ReadAll(resp.Body)
	respString := string(resBody)

	if nil != expectedResBody {
		require.Equal(t, *expectedResBody, respString, "Received unexpected response body")
	}

	return resp.Header.Get("Location")

}

// Checks that the to-do in the system matches expected one
func checkTodo(t *testing.T, todo spi.Todo) {
	response, _ := json.Marshal(todo)
	responseStr := string(response) + "\n"
	testRequest(t, "GET", "/todos/"+todo.Id, "", 200, &responseStr)
}
