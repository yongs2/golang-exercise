package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := initTrie()

	testCases1 := []struct {
		word string
		want string
	}{
		{"sam", "sam."},
		{"john", "john.sam."},
		{"tim", "john.sam.tim."},
		{"jose", "john.se.sam.tim."},
		{"rose", "john.se.rose.sam.tim."},
		{"cat", "cat.john.se.rose.sam.tim."},
		{"dog", "cat.dog.john.se.rose.sam.tim."},
		{"dogg", "cat.dogg..john.se.rose.sam.tim."},
		{"roses", "cat.dogg..john.se.roses..sam.tim."},
	}
	for index, tc := range testCases1 {
		t.Run(fmt.Sprintf("Insert #%d", index), func(t *testing.T) {
			trie.insert(tc.word)
			if got := trie.String(); got != tc.want {
				t.Errorf("got(%s) = %v; want %v", tc.word, got, tc.want)
			}
		})
	}

	testCases2 := []struct {
		word string
		want bool
	}{
		{"sam", true},
		{"john", true},
		{"tim", true},
		{"jose", true},
		{"rose", true},
		{"cat", true},
		{"dog", true},
		{"dogg", true},
		{"roses", true},
		{"rosess", false},
		{"ans", false},
		{"san", false},
	}
	for index, tc := range testCases2 {
		t.Run(fmt.Sprintf("Find #%d", index), func(t *testing.T) {
			if got := trie.find(tc.word); got != tc.want {
				t.Errorf("got(%s) = %v; want %v", tc.word, got, tc.want)
			}
		})
	}
}
