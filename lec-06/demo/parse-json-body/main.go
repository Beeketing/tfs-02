package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create mux router
	router := mux.NewRouter().StrictSlash(true)

	// register handler to router
	router.Methods(http.MethodPost).Path("/welcome").HandlerFunc(welcomeHandler)

	// serve router on port
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// http handler
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	req := Person{}

	// decode json body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprintf(w, "error when parse body")
		return
	}

	fmt.Fprintf(w, "welcome %v, %v year old", req.Name, req.Age)
}
