package main

import (
	"flag"
	"fmt"
)

// go run . -env prod -consumer true
func main() {
	env := flag.String("env", "dev", "Environment(dev, qa, stg, prod)")
	cron := flag.Bool("consumer", false, "boolean")

	flag.Parse()
	fmt.Println("The environment set is", *env)
	fmt.Println("The consumer flag retrieved from command line is", *cron)
}
