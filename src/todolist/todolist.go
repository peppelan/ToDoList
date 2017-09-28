// Provides a simple REST API implementing a basic to-do list
package main

import (
	"expvar"
	"log"
	"net/http"

	"sync"
	"todolist/repo"
	"todolist/spi"
)

var repository spi.Repo

// Starts two HTTP services:
//	one at port 8080 for exposing the ToDoList REST service,
//	one at port 8081 for exposing the expvar service.
// The application runs as long as both HTTP services are up
func main() {
	repository = repo.NewRepo()
	err := repository.Init()

	if nil != err {
		panic(err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go registerBusinessServer(wg)
	go registerExpvarServer(wg)
	wg.Wait()
}

func registerBusinessServer(wg *sync.WaitGroup) {
	log.Fatal(http.ListenAndServe(":8080", NewRouter()))
	wg.Done()
}

func registerExpvarServer(wg *sync.WaitGroup) {
	log.Fatal(http.ListenAndServe(":8081", expvar.Handler()))
	wg.Done()
}
