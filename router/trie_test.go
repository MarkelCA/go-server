package router

import (
    "net/http"
	"testing"
)

func TestInsert(t *testing.T) {
    trie := NewTrieRoutes()
    handler := func(w http.ResponseWriter, r *http.Request) {}

    trie.Insert("apple",mGET, handler)

    if !trie.Search("apple") {
        t.Error("Expected 'apple' to be found in trie.")
    }
}

func TestSearch(t *testing.T) {
    trie := NewTrieRoutes()
    handler := func(w http.ResponseWriter, r *http.Request) {}

    trie.Insert("apple",mGET, handler)

    if !trie.Search("apple") {
        t.Error("Expected 'apple' to be found in trie.")
    }

    if trie.Search("app") {
        t.Error("Expected 'app' to NOT be found in trie.")
    }
}

func TestStartsWith(t *testing.T) {
    trie := NewTrieRoutes()
    handler := func(w http.ResponseWriter, r *http.Request) {}

    trie.Insert("apple",mGET, handler)
    trie.Insert("application",mGET, handler)

    if !trie.StartsWith("app") {
        t.Error("Expected trie to start with 'app'.")
    }

    if trie.StartsWith("banana") {
        t.Error("Expected trie to NOT start with 'banana'.")
    }
}

func TestGetRouteHandler(t *testing.T) {
    trie := NewTrieRoutes()
    handler := func(w http.ResponseWriter, r *http.Request) {}

    trie.Insert("apple",mGET, handler)
    trie.Insert("apple",mPOST, handler)

    if trie.GetRouteHandler("apple", mGET) == nil {
        t.Error("Expected 'apple' GET handler to not be nil.")
    }

    if trie.GetRouteHandler("apple", mPOST) == nil {
        t.Error("Expected 'apple' POST handler to not be nil.")
    }
}
