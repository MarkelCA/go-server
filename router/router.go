package router

import (
    "fmt"
    "net/http"
)

type HttpMethod string

type Handler func(http.ResponseWriter, *http.Request)

const (
    GET  HttpMethod = http.MethodGet
    POST HttpMethod = http.MethodPost
    PUT  HttpMethod = http.MethodPut
)

type Router struct {
    routes []*Route
}

func NewRouter() Router {
    return Router{}
}


func (r Router) GetRoutes() []*Route {
    return r.routes
}

func (r *Router) Add(route Route) {
    r.routes = append(r.routes, &route)
}


func (r *Router) Get(path string, handler Handler) {
    r.Add( Route{path, GET, handler} )
}

func (router Router) Init() {
    for _,r := range router.routes {
        fmt.Printf("Route %v\n", r)
        http.HandleFunc(r.Path, r.Handler)
    }
}



type Route struct {
    Path    string
    Method  HttpMethod
    Handler Handler
}

func NewRoute(path string, method HttpMethod, handler Handler) Route {
    return Route{path, method, handler}
}
