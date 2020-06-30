package main

import (
	"fmt"
	"net/http"
)

// TodoServer currently returns 'TODO ONE' given _any_ request.
func TodoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "TODO ONE")
}
