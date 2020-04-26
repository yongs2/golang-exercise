package trie

import (
	"fmt"
)

const (
	ALPHABET_SIZE = 26
)

type trieNode struct {
	childrens [ALPHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) find(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	if current.isWordEnd {
		return true
	}
	return false
}

func (t *trie) String() string {
	var result string

	current := t.root
	result += StringNode(current)
	return result
}

func StringNode(current *trieNode) string {
	var result string

	for i := 0; i < ALPHABET_SIZE; i++ {
		if current.childrens[i] != nil {
			result += fmt.Sprintf("%c", i+'a')
			result += StringNode(current.childrens[i])
		}
	}
	if current.isWordEnd {
		result += "."
	}
	return result
}
