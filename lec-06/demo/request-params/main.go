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
	router.Methods(http.MethodGet).Path("/hi/{id:(?:\\d+)}").HandlerFunc(hiHandler)

	// serve router on port
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Printf("vars %v", vars)

	name := vars["name"]

	fmt.Fprintf(w, "hi %s", name)
}
