// Provides a simple REST API implementing a basic to-do list
package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"sync"
	"encoding/json"
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
	router.HandleFunc("/todos", todoIndex)
	router.HandleFunc("/todos/{todoId}", todoShow)
	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}