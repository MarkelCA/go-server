package router

import (
    "net/http"
	"testing"
)

func TestInsert(t *testing.T) {
    trie := NewTrie()
    handler := func(w http.ResponseWriter, r *http.Request) {}

    trie.Insert("apple", handler)

    if !trie.Search("apple") {
        t.Error("Expected 'apple' to be found in trie.")
    }
}

func TestSearch(t *testing.T) {
    trie := NewTrie()
    handler := func(w http.ResponseWriter, r *http.Request) {}

    trie.Insert("apple", handler)

    if !trie.Search("apple") {
        t.Error("Expected 'apple' to be found in trie.")
    }

    if trie.Search("app") {
        t.Error("Expected 'app' to NOT be found in trie.")
    }
}

func TestStartsWith(t *testing.T) {
    trie := NewTrie()
    handler := func(w http.ResponseWriter, r *http.Request) {}

    trie.Insert("apple", handler)
    trie.Insert("application", handler)

    if !trie.StartsWith("app") {
        t.Error("Expected trie to start with 'app'.")
    }

    if trie.StartsWith("banana") {
        t.Error("Expected trie to NOT start with 'banana'.")
    }
}

func TestGetHandler(t *testing.T) {
    trie := NewTrie()
    handler := func(w http.ResponseWriter, r *http.Request) {}

    trie.Insert("apple", handler)

    if trie.GetHandler("apple") == nil {
        t.Error("Expected 'apple' handler to not be nil.")
    }
}
