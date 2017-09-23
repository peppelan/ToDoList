// Provides a simple REST API implementing a basic to-do list
package todolist

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", Handler()))
}

func Handler() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
