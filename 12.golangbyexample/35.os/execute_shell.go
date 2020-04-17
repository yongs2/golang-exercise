package main

import (
	"log"
	"os/exec"
)

/*
echo '#!/bin/bash' > sample.sh
echo 'echo "Triggered from .go file"'>> sample.sh
chmod +x sample.sh
*/
func main() {
	out, err := exec.Command("/bin/sh", "sample.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
