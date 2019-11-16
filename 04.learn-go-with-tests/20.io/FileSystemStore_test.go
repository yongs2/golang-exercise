package main

import (
	"testing"
	"strings"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/leagure from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name":"Cleo", "Wins":10},
			{"Name":"Chris", "Wins":33}]`)

		store := FileSystemStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
	})
}
