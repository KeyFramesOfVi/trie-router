package types

import (
	"errors"
	"strings"
)

type node struct {
	wildcard bool
	children map[string]*node
	value    interface{}
}

func newNode() *node {
	return &node{
		wildcard: false,
		children: make(map[string]*node),
		value:    nil,
	}
}

// Trie : Prefix Tree where each node is part of the address
// for a value
type Trie struct {
	Delimiter string
	Root      *node
}

func NewTrie() *Trie {
	return &Trie{
		Delimiter: defaultDelimiter,
		Root:      newNode(),
	}
}

const (
	defaultDelimiter = "/"
)

func (root *Trie) splitPattern(pattern string) []string {
	p := strings.TrimRight(pattern, root.Delimiter)
	pathArray := strings.Split(p, root.Delimiter)
	if pathArray[0] == "" {
		pathArray[0] = string(pattern[0])
	}
	return pathArray
}

// Put establishes a value for a given pattern
func (root *Trie) Put(pattern string, value interface{}) error {
	pathArray := root.splitPattern(pattern)
	curr := root.Root
	if len(pattern) == 0 {
		curr.children["/"].value = value
		return nil
	}

	for _, path := range pathArray {
		trimmedPath := strings.TrimSpace(path)
		_, itExists := curr.children[trimmedPath]
		if itExists == false {
			curr.children[path] = newNode()
		}
		curr = curr.children[trimmedPath]

	}
	curr.value = value
	return nil
}

// Get returns the value associated with the pattern given
func (root *Trie) Get(pattern string) (interface{}, error) {
	pathArray := root.splitPattern(pattern)
	curr := root.Root
	if len(pattern) == 0 {
		value, itExists := curr.children["/"]
		if itExists {
			return value, nil
		}
		return nil, errors.New("Error: Pattern not found")
	}

	for _, path := range pathArray {
		trimmedPath := strings.TrimSpace(path)
		_, itExists := curr.children[trimmedPath]
		if itExists {
			curr = curr.children[trimmedPath]
		} else {
			return nil, errors.New("Error: Pattern not found")
		}
	}
	return curr.value, nil
}
