package main

import (
	"fmt"
	"log"
	"sync"
)

type customQueue struct {
	queue []string
	lock  sync.RWMutex
}

func (c *customQueue) Enqueue(val string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.queue = append(c.queue, val)
}

func (c *customQueue) Dequeue() error {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.queue = c.queue[1:]
		return nil
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.queue[0], nil
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return len(c.queue)
}

func (c *customQueue) Empty() bool {
	return len(c.queue) == 0
}

func main() {
	customQueue := &customQueue{
		queue: make([]string, 0),
	}
	log.Printf("Enqueue: A\n")
	customQueue.Enqueue("A")
	log.Printf("Enqueue: B\n")
	customQueue.Enqueue("B")
	log.Printf("Size: %d\n", customQueue.Size())
	for customQueue.Size() > 0 {
		frontVal, _ := customQueue.Front()
		log.Printf("Front: %s\n", frontVal)
		log.Printf("Dequeue: %s\n", frontVal)
		customQueue.Dequeue()
	}
	log.Printf("Size: %d\n", customQueue.Size())
}
