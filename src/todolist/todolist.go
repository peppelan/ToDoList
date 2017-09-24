// Provides a simple REST API implementing a basic to-do list
package main

import (
	"expvar"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"sync"
)

// Starts two HTTP services:
//	one at port 8080 for exposing the ToDoList REST service,
//	one at port 8081 for exposing the expvar service.
// The application runs as long as both HTTP services are up
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go registerBusinessServer(wg);
	go registerExpvarServer(wg);
	wg.Wait()
}

func registerBusinessServer(wg *sync.WaitGroup) {
	log.Fatal(http.ListenAndServe(":8080", handler()))
	wg.Done()
}

func registerExpvarServer(wg *sync.WaitGroup) {
	log.Fatal(http.ListenAndServe(":8081", expvar.Handler()))
	wg.Done()
}

func handler() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
