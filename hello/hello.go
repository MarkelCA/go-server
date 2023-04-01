package hello

import (
	"fmt"
	"net/http"
)

func HelloGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Get Client!")
}

func MeGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Me!")
}

func HelloPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Post Client!")
}

func HelloUsertHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%v", r)
}
