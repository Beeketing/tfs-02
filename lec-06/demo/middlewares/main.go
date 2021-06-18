package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create mux router
	router := mux.NewRouter().StrictSlash(true)

	// register handler to router
	router.Methods(http.MethodGet).Path("/hello").
		Handler(authenticateMiddleware(&HelloHandler{}))

	// using middleware function
	router.Use(contentTypeCheckingMiddleware)

	// serve router on port
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

// http handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

const JsonContentType = "application/json"

// middleware func
func contentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != JsonContentType {
			fmt.Fprintf(w, "request only allow content type application/json")
			return
		}

		fmt.Fprintf(w, "hello again")

		// main handler
		next.ServeHTTP(w, r)
	})
}

func authenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Authen")

		if a != "123456" {
			fmt.Fprintf(w, "invalid authen")
			return
		}

		next.ServeHTTP(w, r)
	})
}
