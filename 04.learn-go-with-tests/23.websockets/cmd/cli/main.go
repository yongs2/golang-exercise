package main

import (
	"fmt"
	poker "04.learn-go-with-tests/23.websockets"
	"os"
	"log"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store )

	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
