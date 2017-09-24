package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


// Creates a router for the ToDoList app.
// The router is a github.com/gorilla/mux, with handlers for the ToDoList application,
// and instrumented for logging relevant information about the HTTP requests and their
// handling.
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(logger(route.HandlerFunc, route.Name))
	}

	return router
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s\t%d us",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			name,
			time.Since(start).Nanoseconds()/1e3,
		)
	})
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		todoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		todoShow,
	},
}
