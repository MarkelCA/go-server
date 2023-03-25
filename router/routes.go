package router

import (
    "net/http"
    "net/url"
    "fmt"
)

type httpMethod string

const (
	mGET httpMethod = http.MethodGet
	mPOST           = http.MethodPost
	mPUT            = http.MethodPut
	mDELETE         = http.MethodDelete
)

var strToMethod = map[string]httpMethod{
    http.MethodGet    : mGET,
    http.MethodDelete : mDELETE,
    http.MethodPost   : mPOST,
    http.MethodPut    : mPUT,
}


type Handlers map[httpMethod]http.HandlerFunc
type Routes   map[url.URL]Handlers

func NewRoutes() Routes {
    return Routes{}
}

func (r *Routes) Get(path string, handler http.HandlerFunc) {
    u := url.URL{
        Path: path,
    }
    r.addRoute(u, mGET, handler)
}

func (r *Routes) Post(path string, handler http.HandlerFunc) {
    u := url.URL{
        Path: path,
    }
    r.addRoute(u, mPOST, handler)
}

func (r *Routes) addRoute(path url.URL, method httpMethod, handler http.HandlerFunc) {
    if _, pathExists := (*r)[path] ; pathExists {
        (*r)[path][method] = handler
    } else {
        (*r)[path] = Handlers{
            method : handler,
        }
    }

    fmt.Printf("Added route %-4v -> %v\n", method, path.String()) // Method right-padded with 4 spaces

}


// Gets the global handler function.
// This function acts as the handler for all the requests.
// Firstly checks that the route exists and that the method
// is allowed, then maps the request to the specific handler 
// function defined in the router map.
func (router Routes) getHandler() http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request) {
        if _,pathExists := router[*r.URL] ; pathExists == false {
            http.Error(w, "404 Not Found", http.StatusNotFound)
            return
        } 

        method := strToMethod[r.Method]
        if _,methodAllowed := router[*r.URL][method]; methodAllowed == false {
            w.Header().Set("Allow", r.Method)
            http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
            return
        }

        router[*r.URL][method](w,r)
    }
}


