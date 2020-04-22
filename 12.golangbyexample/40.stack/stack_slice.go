package main

import (
	"fmt"
	"log"
	"sync"
)

type customStack struct {
	stack []string
	lock  sync.RWMutex
}

func (c *customStack) Push(value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, value)
}

func (c *customStack) Pop() error {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len-1]
		return nil
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customStack) Front() (string, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[len-1], nil
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customStack) Size() int {
	return len(c.stack)
}

func (c *customStack) Empty() bool {
	return len(c.stack) == 0
}

func main() {
	customQueue := &customStack{
		stack: make([]string, 0),
	}
	log.Printf("Push: %s\n", "A")
	customQueue.Push("A")

	log.Printf("Push: %s\n", "B")
	customQueue.Push("B")

	log.Printf("Size: %d\n", customQueue.Size())
	for customQueue.Size() > 0 {
		frontVal, _ := customQueue.Front()
		log.Printf("Front: %s\n", frontVal)
		log.Printf("Pop: %s\n", frontVal)
		customQueue.Pop()
	}
	log.Printf("Size: %d\n", customQueue.Size())
}
