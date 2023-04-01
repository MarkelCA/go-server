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
}

func NewRouter() *Mux {
    routes := &Routes{}
    return &Mux{
        routes: routes,
        handler : routes.getHandler(),
    }
}

// Default implementation for the Router interface
// Contains the routes map and the main handler 
// function, as well as the http verbs functions
// to add more handlers.
type Mux struct {
    routes *Routes
    handler http.HandlerFunc
}

func (m *Mux) Get(path string, handler http.HandlerFunc) {
    m.routes.Get(path, handler)
}

func (m *Mux) Post(path string, handler http.HandlerFunc) {
    m.routes.Post(path, handler)
}

func (m Mux) Routes() *Routes{
    return m.routes
}


func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handling request")

    r.URL.Path = removeTrailingSlash(r.URL.String())
    m.handler(w, r)
}

