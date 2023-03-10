package main


import (
    "fmt"
    "log"
    "net/http"
    "github.com/markelca/go-server/hello"
)


func main() {
    hello.HelloServer()
    http.HandleFunc("/hello", hello.HelloHandler) // Update this line of code

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
