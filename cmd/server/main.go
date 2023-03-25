package main

import (
	"fmt"
    "errors"
	"log"
	"net/http"
	"github.com/markelca/go-server/hello"
	"github.com/markelca/go-server/router"
)

func main() {
    r := router.NewRouter()

    r.Get("/hello", hello.HelloGetHandler)
    r.Post("/hello", hello.HelloPostHandler)
    r.Get("/hello/me", hello.MeGetHandler)


    err := http.ListenAndServe(":8080", r)

    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Fatal("error starting server: %s\n", err)
	}

}
