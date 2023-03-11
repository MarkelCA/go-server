package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/markelca/go-server/hello"
	"github.com/markelca/go-server/router"
)

func main() {
    router := router.NewRouter()
    router.Get("/hello", hello.HelloHandler)

    router.Init()

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
