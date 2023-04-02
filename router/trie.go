package router

import (
    "fmt"
    "net/http"
)

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

func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, c := range word {
		if _, ok := currentNode.children[c]; !ok {
			currentNode.children[c] = &node{
				children: make(map[rune]*node),
				isEnd:    false,
                handler:  nil,
			}
		}
		currentNode = currentNode.children[c]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, c := range word {
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

func (t *Trie) Print() {
    t.printHelper(t.root, []rune{})
}

func (t *Trie) printHelper(node *node, word []rune) {
    if node.isEnd {
        fmt.Println(string(word))
    }
    for ch, child := range node.children {
        t.printHelper(child, append(word, ch))
    }
}

func main () {
    trie := NewTrie()
    trie.Insert("apple")
    trie.Insert("application")

    trie.Print()

    fmt.Println(trie.Search("app"))
    fmt.Println(trie.StartsWith("app"))
}


