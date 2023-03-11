package router

import (
    "fmt"
    "net/http"
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

func allowMethod(h http.HandlerFunc, method string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if method != r.Method {
            w.Header().Set("Allow", method)
            http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
            return
        }
        h(w, r)
    }
}

func (r *Router) Get(path string, rawHandler http.HandlerFunc) {
    handler := allowMethod(rawHandler, http.MethodGet)
    r.Add( Route{path, http.MethodGet, handler } )
}

func (r *Router) Post(path string, rawHandler http.HandlerFunc) {
    handler := allowMethod(rawHandler, http.MethodPost)
    r.Add( Route{path, http.MethodPost, handler } )
}

func (router Router) Init(mux *http.ServeMux) {
    for _,r := range router.routes {
        fmt.Printf("Added route %v (%v)\n", r.Path, r.Method)
        mux.HandleFunc(r.Path, r.Handler)
    }
}


type Route struct {
    Path    string
    Method  string
    Handler http.HandlerFunc
}

func NewRoute(path string, method string, handler http.HandlerFunc) Route {
    return Route{path, method, handler}
}
