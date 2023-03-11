package router

import (
    "fmt"
    "net/http"
)

type Handler struct {
    Method  string
    Handler http.HandlerFunc
}
type Router map[string]Handler

func NewRouter() Router {
    return Router{}
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
    fmt.Println(handler)
    (*r)[path] = Handler{
        http.MethodGet, 
        handler,
    }
}

func (r *Router) Post(path string, rawHandler http.HandlerFunc) {
    handler := allowMethod(rawHandler, http.MethodPost)
    fmt.Println(handler)
}

func (router Router) Init(mux *http.ServeMux) {
    for path,r := range router {
        fmt.Printf("Added route %v -- %v\n", path, r)
        //fmt.Printf("Added route %v (%v)\n", r.Path, r.Method)
        mux.HandleFunc(path, r.Handler)
    }
}

