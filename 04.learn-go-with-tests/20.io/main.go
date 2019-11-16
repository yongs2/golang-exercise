package main

import (
	"log"
	"net/http"
)

// go run .
// curl -X POST http://localhost:5000/players/Pepper
// curl http://localhost:5000/players/Pepper
func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
