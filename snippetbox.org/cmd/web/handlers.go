package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home ...
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("om gam gan ganapatye namo namah"))
}

// ShowSnippet ...
func ShowSnippet(w http.ResponseWriter, r *http.Request) {

	// Grab ID from GET and sanitize to int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// 404 If Wasn't Int, or Less Than 1
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display snippet id %d", id)
}

// NewSnippet ...
func NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Render the New Snippet form"))
}
