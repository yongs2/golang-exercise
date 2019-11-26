package main

import (
	"log"
	"net/http"
	"04.learn-go-with-tests/23.websockets"
)

// go run main.go
// curl -X POST http://localhost:5000/players/Pepper
// curl http://localhost:5000/players/Pepper

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer close()
	
	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
