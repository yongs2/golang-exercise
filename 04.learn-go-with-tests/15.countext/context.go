package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	fetch := store.Fetch()
	fmt.Println("Server:", fetch)
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprintf(w, fetch)
	}
}
