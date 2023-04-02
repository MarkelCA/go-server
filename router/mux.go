package router

import (
    "fmt"
    "net/http"
)

// Multiplexing Router interface
// It describes the common functions that any
// implementation for the http request multiplexers
// must have.
type Router interface {
    http.Handler
    Routes
    Get(path string, handler http.HandlerFunc)
    Post(path string, handler http.HandlerFunc)
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// It describes the common functions for a Routes object.
// Useful to change the route handling/searching 
// implementation.
// e.g: HashMap based vs Trie based etc.
type Routes interface {
    getHandler() http.HandlerFunc
    add(path string, method httpMethod, handler http.HandlerFunc)
}

func NewRouter(routes Routes) *Mux {
    return &Mux{
        Routes: routes,
        handler : routes.getHandler(),
    }
}

// Default implementation for the Router interface
// Contains the routes map and the main handler 
// function, as well as the http verbs functions
// to add more handlers.
type Mux struct {
    Routes Routes
    handler http.HandlerFunc
}

func (m *Mux) Get(path string, handler http.HandlerFunc) {
    m.Routes.add(path, mGET, handler)
}

func (m *Mux) Post(path string, handler http.HandlerFunc) {
    m.Routes.add(path, mPOST, handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handling request")

    r.URL.Path = removeTrailingSlash(r.URL.String())
    m.handler(w, r)
}

