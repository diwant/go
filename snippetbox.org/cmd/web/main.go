package main

import (
	"log"
	"net/http"
)

func main() {

	// Set up a Serve Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	// Set up a Server
	log.Println("Serving on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
