package router

import (
    "fmt"
    "net/http"
)

type Handlers map[string]http.HandlerFunc

//type route string
type Router map[string]Handlers

func NewRouter() Router {
    return Router{}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
    r.addRoute(path, http.MethodGet, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
    r.addRoute(path, http.MethodPost, handler)
}

func (r *Router) addRoute(path string, method string, handler http.HandlerFunc) {
    if _, pathExists := (*r)[path] ; pathExists {
        (*r)[path][method] = handler
    } else {
        (*r)[path] = Handlers{
            method : handler,
        }
    }

    fmt.Printf("Added route %v (%v) -> %v\n", path, method, handler)

}


func (r Router) Init(mux *http.ServeMux) {
    for path,handlers := range r {

        for _,h := range handlers {
            //fmt.Printf("%v:: %v\n",m, h)
            h2 := r.mergeHandlers(h)
            mux.HandleFunc(path, h2)
            break
        }
    }
}

func (router Router) mergeHandlers(h http.HandlerFunc) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request) {
        if _,pathExists := router[r.URL.String()] ; pathExists == false {
            http.Error(w, "404 method not allowed", http.StatusNotFound)
            return
        } 
        if _,methodAllowed := router[r.URL.String()][r.Method]; methodAllowed == false {
            w.Header().Set("Allow", r.Method)
            http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
            return
        }

        router[r.URL.String()][r.Method](w,r)
    }
}


