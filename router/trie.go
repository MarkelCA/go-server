package router

import (
    "fmt"
    "net/http"
)

// Routes interface implementation using a Radix
// trie data structure
type TrieRoutes Trie

func NewTrieRoutes() TrieRoutes {
    trie := NewTrie()
    return TrieRoutes{trie.root}
}

type Trie struct {
	root *node
}

type node struct {
	children map[rune]*node
	isEnd    bool
    handler  *http.HandlerFunc
}

func NewTrie() *Trie {
	return &Trie{
		root: &node{
			children: make(map[rune]*node),
			isEnd:    false,
            handler:  nil,
		},
	}
}

func (t *Trie) Insert(route string, h http.HandlerFunc) {
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

func (t *Trie) Search(route string) bool {
	currentNode := t.root
	for _, c := range route {
		if _, ok := currentNode.children[c]; !ok {
			return false
		}
		currentNode = currentNode.children[c]
	}
	return currentNode.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t.root
	for _, c := range prefix {
		if _, ok := currentNode.children[c]; !ok {
			return false
		}
		currentNode = currentNode.children[c]
	}
	return true
}

func (t Trie) GetHandler(route string) *http.HandlerFunc {
	currentNode := t.root
	for _, c := range route {
		if _, ok := currentNode.children[c]; !ok {
			return nil
		}
		currentNode = currentNode.children[c]
	}
	return currentNode.handler
}

func (t *Trie) Print() {
    t.printHelper(t.root, []rune{})
}

func (t *Trie) printHelper(node *node, route []rune) {
    if node.isEnd {
        fmt.Println(string(route))
    }
    for ch, child := range node.children {
        t.printHelper(child, append(route, ch))
    }
}

func main () {
    trie := NewTrie()
    handler := func(w http.ResponseWriter, r *http.Request) {}
    trie.Insert("apple", handler)
    trie.Insert("application", handler)

    trie.Print()

    fmt.Println(trie.Search("app"))
    fmt.Println(trie.StartsWith("app"))
}

