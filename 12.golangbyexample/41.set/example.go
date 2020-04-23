package main

import (
	"fmt"
	"log"
)

func makeSet() *customSet {
	return &customSet{
		container: make(map[string]struct{}),
	}
}

type customSet struct {
	container map[string]struct{}
}

func (c *customSet) Exists(key string) bool {
	_, exists := c.container[key]
	return exists
}

func (c *customSet) Add(key string) {
	c.container[key] = struct{}{}
}

func (c *customSet) Remove(key string) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *customSet) Size() int {
	return len(c.container)
}

func main() {
	var key string

	customSet := makeSet()

	key = "A"
	log.Printf("Add: %s\n", key)
	customSet.Add(key)

	key = "B"
	log.Printf("Add: %s\n", key)
	customSet.Add(key)

	log.Printf("Size: %d, customSet=[%+v]\n", customSet.Size(), customSet)
	key = "A"
	log.Printf("%s Exists?: %t\n", key, customSet.Exists(key))
	key = "B"
	log.Printf("%s Exists?: %t\n", key, customSet.Exists(key))
	key = "C"
	log.Printf("%s Exists?: %t\n", key, customSet.Exists(key))

	key = "B"
	log.Printf("Remove: %s\n", key)
	customSet.Remove(key)
	log.Printf("%s Exists?: %t\n", key, customSet.Exists(key))
}
