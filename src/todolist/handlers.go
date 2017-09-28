package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"html/template"
	"io"
	"io/ioutil"
	"todolist/spi"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// Responds 200 (OK)
func todoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repository.FindAll()); err != nil {
		panic(err)
	}
}

// Responds 200 (OK) or 404 (Not Found) for non-stored IDs
func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	todo := repository.Find(todoId)

	if nil == todo {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	}
}

// Responds 201 (Created) when successful,
// 422 (Unprocessable entity) when the provided object does not correctly translate to a to-do,
// or 406 (Not acceptable) when an ID was provided - the application is responsible for creating it
func todoCreate(w http.ResponseWriter, r *http.Request) {
	var todo spi.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	id := repository.Create(todo)
	w.Header().Set("Location", "/todos/"+template.URLQueryEscaper(id))
	w.WriteHeader(http.StatusCreated)
}

// Responds 200 (OK) when deleted, or 404 (Not Found) for non-stored IDs
func todoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	found := repository.Destroy(todoId)

	if found {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

// Responds 200 (OK) when updated, 404 (Not Found) for non-stored IDs,
// 422 (Unprocessable entity) when the provided object does not correctly translate to a to-do
func todoUpdate(w http.ResponseWriter, r *http.Request) {

	// Get the to-do ID to update
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Read the updated to-do
	var todo spi.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	updated := repository.Update(todoId, todo)
	if updated {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
