package mux

import (
    "fmt"
    "net/http"
)

type OtherHandler struct {
}

func (h *OtherHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {

    fmt.Fprintf(w, "Hello there!")
}
