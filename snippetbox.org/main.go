package main

import (
	"fmt"
	"log"
	"net/http"
)

// Home ...
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("om gam gan ganapatye namo namah"))
}

func main() {
	fmt.Println("gam")

	// New ServeMux (Router)
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	// Start the Server
	log.Println("Starting server")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
