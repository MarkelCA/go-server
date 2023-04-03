package router

import (
    "fmt"
    "net/http"
)

// Routes interface implementation using a Radix
// trie data structure
//type TrieRoutes TrieRoutes
type TrieRoutes struct {
	root *node
}

type node struct {
	children map[rune]*node
	isEnd    bool
    handler  *http.HandlerFunc
}

func NewTrieRoutes() *TrieRoutes {
	return &TrieRoutes{
		root: &node{
			children: make(map[rune]*node),
			isEnd:    false,
            handler:  nil,
		},
	}
}

func (t *TrieRoutes) add(route string, m httpMethod, h http.HandlerFunc) {
    t.Insert(route, h)
}

func (t *TrieRoutes) Insert(route string, h http.HandlerFunc) {
	currentNode := t.root
	for _, c := range route {
		if _, ok := currentNode.children[c]; !ok {
			currentNode.children[c] = &node{
				children: make(map[rune]*node),
				isEnd:    false,
                handler:  &h,
			}
		}
		currentNode = currentNode.children[c]
	}
	currentNode.isEnd = true
}

func (t *TrieRoutes) Search(route string) bool {
	currentNode := t.root
	for _, c := range route {
		if _, ok := currentNode.children[c]; !ok {
			return false
		}
		currentNode = currentNode.children[c]
	}
	return currentNode.isEnd
}

func (t *TrieRoutes) StartsWith(prefix string) bool {
	currentNode := t.root
	for _, c := range prefix {
		if _, ok := currentNode.children[c]; !ok {
			return false
		}
		currentNode = currentNode.children[c]
	}
	return true
}

func (t TrieRoutes) GetRouteHandler(route string) http.HandlerFunc {
	currentNode := t.root
	for _, c := range route {
		if _, ok := currentNode.children[c]; !ok {
			return nil
		}
		currentNode = currentNode.children[c]
	}
	return *currentNode.handler
}

func (t *TrieRoutes) Print() {
    t.printHelper(t.root, []rune{})
}

func (t *TrieRoutes) printHelper(node *node, route []rune) {
    if node.isEnd {
        fmt.Println(string(route))
    }
    for ch, child := range node.children {
        t.printHelper(child, append(route, ch))
    }
}

func (router TrieRoutes) GetHandler() http.HandlerFunc{
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

        //method := strToMethod[r.Method]
        path   := r.URL.Path
        router.GetRouteHandler(path)(w,r)
        //router.handle(w,r)
    }
}
