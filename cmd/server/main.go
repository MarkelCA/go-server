package main

import (
	"fmt"
    "errors"
	"log"
	"net/http"
	//"net/url"
	"github.com/markelca/go-server/hello"
	"github.com/markelca/go-server/router"
)

func main() {
    //routes := router.NewMapRoutes()
    routes := router.NewTrieRoutes()
    r := router.NewRouter(routes)

    r.Get("/hello", hello.HelloGetHandler)
    r.Post("/hello", hello.HelloPostHandler)
    r.Get("/hello/me", hello.MeGetHandler)

    r.Get("user/{id}", hello.HelloUsertHandler)

    router.PrintRoutes(r.Routes)

    err := http.ListenAndServe(":8080", r)

    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Fatal("error starting server: %s\n", err)
	}

}
