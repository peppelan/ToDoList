package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Describes how the endpoints handled by the app are wired
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
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		todoCreate,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/todos/{todoId}",
		todoDelete,
	},
	Route{
		"TodoUpdate",
		"PUT",
		"/todos/{todoId}",
		todoUpdate,
	},
}
