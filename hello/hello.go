package hello

import (
	"fmt"
	"net/http"
)

func HelloGetHandler(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/hello" {
		//http.Error(w, "404 not found.", http.StatusNotFound)
		//return

	fmt.Fprintf(w, "Hello Get Client!")
}

func HelloPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Post Client!")
}
