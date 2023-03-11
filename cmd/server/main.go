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
    mux := http.NewServeMux()
    router := router.NewRouter()
    router.Get("/hello", hello.HelloGetHandler)
    router.Post("/hello2", hello.HelloPostHandler)

    router.Init(mux)

	fmt.Printf("Starting server at port 8080\n")

    err := http.ListenAndServe(":8080", mux)

    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Fatal("error starting server: %s\n", err)
	}

}
