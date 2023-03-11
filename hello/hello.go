package hello

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method "+r.Method+" is not supported for this request.", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello Client!")
}
