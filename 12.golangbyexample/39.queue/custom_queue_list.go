package main

import (
	"container/list"
	"fmt"
	"log"
)

type customQueue struct {
	queue *list.List
}

func (c *customQueue) Enqueue(val string) {
	c.queue.PushBack(val)
}

func (c *customQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		c.queue.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return c.queue.Len()
}

func (c *customQueue) Empty() bool {
	return c.queue.Len() == 0
}

func main() {
	customQueue := &customQueue{
		queue: list.New(),
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
