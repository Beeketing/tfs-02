package handlers

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello guys")
}
