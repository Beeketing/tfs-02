package server

import (
	"fmt"
	"net/http"

	"../handlers" // import local package
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8090/hello")
	// Defer function will be called when process exits
	defer func() {
		fmt.Println("Server is stopped")
	}()
	//////////////////////////////////////////////////////////////////////////////////
	// register handlers
	// --> browse to http://localhost:8090/hello will return Hello guys
	http.HandleFunc("/hello", handlers.Hello) // simple hello
	// hi handler will read query param
	// --> browse to http://localhost:8090/hi?name=Trung. will return Hi Trung
	http.HandleFunc("/hi", handlers.HiWithParam)
	//////////////////////////////////////////////////////////////////////////////////
	// run server
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic("Error when running server")
	}
}
