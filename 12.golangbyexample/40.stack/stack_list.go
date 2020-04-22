package main

import (
	"container/list"
	"fmt"
	"log"
)

type customStack struct {
	stack *list.List
}

func (c *customStack) Push(value string) {
	c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customStack) Front() (string, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customStack) Size() int {
	return c.stack.Len()
}

func (c *customStack) Empty() bool {
	return c.stack.Len() == 0
}

func main() {
	customQueue := &customStack{
		stack: list.New(),
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
