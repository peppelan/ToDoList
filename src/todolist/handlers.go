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

// Responds 200 (OK) or TODO 404 (Not Found) for non-stored IDs
func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repository.Find(todoId)); err != nil {
		panic(err)
	}
}

// Responds 201 (Created) or TODO 409 (Conflict) for already-stored IDs
func todoCreate(w http.ResponseWriter, r *http.Request) {
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
	}

	t := repository.Create(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", "/todos/"+template.URLQueryEscaper(t.Id))
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
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
