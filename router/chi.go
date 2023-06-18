package router

import (
    "github.com/go-chi/chi/v5"
    "net/http"
    "fmt"
)

func NewChiRouter() Router {
    router := &ChiRouter{
        mux: chi.NewRouter(),
    }
    return router

}

type ChiRouter struct {
    mux *chi.Mux
}

func (r ChiRouter) PrintRoutes() {
    err := chi.Walk(r.mux, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the routes: %s\n", err.Error())
	}
}

func (r ChiRouter) Get(path string, handler http.HandlerFunc) {
    r.mux.Get(path, handler)
}

func (r ChiRouter) Post(path string, handler http.HandlerFunc) {
    r.mux.Post(path, handler)
}

func (r ChiRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    r.mux.ServeHTTP(w,  req)
}

