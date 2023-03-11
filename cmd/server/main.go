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
    //r := router.NewRoute("/hello", http.MethodGet, hello.HelloGetHandler)
    r := router.Router{
        "/hello" : router.Handler{
            http.MethodGet,
            hello.HelloGetHandler,
        },
    }
    fmt.Println(r)
    //fmt.Println(r)

    theRouter := router.NewRouter()
    fmt.Println(theRouter)
    theRouter.Get("/hello", hello.HelloGetHandler)
    //theRouter.Post("/hello2", hello.HelloPostHandler)

    theRouter.Init(mux)

	//fmt.Printf("Starting server at port 8080\n")

    err := http.ListenAndServe(":8080", mux)

    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Fatal("error starting server: %s\n", err)
	}

}
