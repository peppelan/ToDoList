package main

import (
	"fmt"
	"github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
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
		// Create our HTTP collector.
		collector, err := zipkin.NewHTTPCollector(zipkinUrl + "/api/v1/spans")
		if err != nil {
			fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
			os.Exit(-1)
		}

		// Create our recorder.
		recorder := zipkin.NewRecorder(collector, false, "0.0.0.0:0", "todolist")

		// Create our tracer.
		tracer, err := zipkin.NewTracer(
			recorder,
			zipkin.ClientServerSameSpan(true),
			zipkin.TraceID128Bit(true),
		)
		if err != nil {
			fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
			os.Exit(-1)
		}

		// Explicitly set our tracer to be the default tracer.
		opentracing.InitGlobalTracer(tracer)

		fmt.Printf("Initialized tracer")

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
			opentracing.StartSpan(name)
			span, ctx := opentracing.StartSpanFromContext(r.Context(), name)
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
