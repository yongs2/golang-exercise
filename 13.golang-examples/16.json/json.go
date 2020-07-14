package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var strings []string
	var jsonstring = `["lorem", "ipsum", "dolor", "sit", "amet"]`

	err := json.Unmarshal([]byte(jsonstring), &strings)
	if err != nil {
		log.Println("error while unmarshalling")
		os.Exit(2)
	}
	log.Printf("Unmarshal[%T]=[%v]\n", strings, strings)

	jsonData, err := json.Marshal(strings)
	if err != nil {
		log.Println("error while marshalling")
		os.Exit(2)
	}

	log.Printf("Marshal=[%+v]\n", string(jsonData))
}
