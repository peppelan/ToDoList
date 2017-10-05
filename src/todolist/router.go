package main

import (
	"github.com/gorilla/mux"
	"github.com/ricardo-ch/go-tracing"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	zipkinUrl = os.Getenv("ZIPKIN_URL")
)

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

	if "" != zipkinUrl {
		// Enable zipkin logging
		tracing.SetGlobalTracer("todolist", zipkinUrl)
		defer tracing.FlushCollector()
	}

	return router
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		responseWriter := NewLoggingResponseWriter(w)

		defer func() {
			if err := recover(); err != nil {
				responseWriter.WriteHeader(http.StatusInternalServerError)

				log.Printf(
					"%s\t%s\t%s\t%s\t%d (%s)\t%d us: %d",
					r.RemoteAddr,
					r.Method,
					r.RequestURI,
					name,
					responseWriter.statusCode,
					http.StatusText(responseWriter.statusCode),
					time.Since(start).Nanoseconds()/1e3,
					err,
				)
			}
		}()

		if "" != zipkinUrl {

			// Define a trace
			span, ctx := tracing.CreateSpanFromClientContext(r, name, nil)
			defer span.Finish()

			// Serve the request, and propagate the new context
			inner.ServeHTTP(responseWriter, r.WithContext(ctx))
		} else {
			// Serve the request only
			inner.ServeHTTP(responseWriter, r)
		}

		log.Printf(
			"%s\t%s\t%s\t%s\t%d (%s)\t%d us",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			name,
			responseWriter.statusCode,
			http.StatusText(responseWriter.statusCode),
			time.Since(start).Nanoseconds()/1e3,
		)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
