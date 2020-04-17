package main

import (
	"log"
	"os"
)

func main() {
	err := os.Setenv("a", "b")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Setenv("c", "d")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ENV:", os.Environ())

	val := os.Getenv("a")
	log.Println("a:", val)

	err = os.Unsetenv("a")
	if err != nil {
		log.Fatal(err)
	}

	val, present := os.LookupEnv("a")
	log.Printf("%s env variable present: %t, val: %s\n", "a", present, val)

	val, present = os.LookupEnv("c")
	log.Printf("%s env variable present: %t, val: %s\n", "c", present, val)

	os.Clearenv()

	log.Println("ENV:", os.Environ())
}
