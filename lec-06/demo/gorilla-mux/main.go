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
	router.Methods(http.MethodGet).Path("/hello").HandlerFunc(helloHandler)

	// serve router on port
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}

// http handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
