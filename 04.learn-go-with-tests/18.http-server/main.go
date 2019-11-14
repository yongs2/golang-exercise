package main

import (
	"net/http"
	"log"
)

type InMemoryPlayerStore struct {}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// curl http://localhost:5000/players/Pepper
func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
