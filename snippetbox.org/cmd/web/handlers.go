package main

import "net/http"

// Home ...
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("om gam gan ganapatye namo namah"))
}
