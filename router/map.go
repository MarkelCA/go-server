package router

import (
    "net/http"
    "strings"
    "fmt"
)


// Routes interface implementation using a map
// data structure
type MapRoutes   map[string]Handlers
type Handlers map[httpMethod]http.HandlerFunc

func NewMapRoutes() *MapRoutes {
    return &MapRoutes{}
}

func (r *MapRoutes) add(path string, method httpMethod, handler http.HandlerFunc) {
    removeTrailingSlash(path)
    if _, pathExists := (*r)[path] ; pathExists {
        (*r)[path][method] = handler
    } else {
        (*r)[path] = Handlers{
            method : handler,
        }
    }

    fmt.Printf("Added route %-4v -> %v\n", method, path) // Method right-padded with 4 spaces

}

// Gets the global handler function.
// This function acts as the handler for all the requests.
// Firstly checks that the route exists and that the method
// is allowed, then maps the request to the specific handler 
// function defined in the router map.
func (router MapRoutes) getHandler() http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request) {
        //if router.exists(r.URL.Path) == false {
            //http.Error(w, "404 Not Found", http.StatusNotFound)
            //return
        //} 

        //method := strToMethod[r.Method]
        //if _,methodAllowed := router[r.URL.Path][method]; methodAllowed == false {
            //w.Header().Set("Allow", r.Method)
            //http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
            //return
        //}

        router.handle(w,r)
    }
}

func (router MapRoutes) handle(w http.ResponseWriter, r *http.Request) {
    method := strToMethod[r.Method]
    path   := r.URL.Path
    //fmt.Printf("%v, %v", method, r.URL.Path)
    if handler, handlerExists := router[path][method] ; handlerExists {
        fmt.Println("642642642")
        handler(w,r)
    } else {
        fmt.Println("HIIII")
        fmt.Println(strings.Index(path, "{"))

        for pos,_ := range path {
            currentPath := path[:pos + 1]
            if router.exists(currentPath) {
                fmt.Printf("found %v", currentPath)
            }

            //fmt.Printf("%c -> %v",char, router.exists(path[:pos]))
        }
    }
}

func (router MapRoutes) exists(path string) bool {
    _, result := router[path]
    return result
}


