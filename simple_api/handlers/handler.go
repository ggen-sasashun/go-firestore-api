package handlers

import (
	"fmt"
	"net/http"
)

// GET /hello
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}
