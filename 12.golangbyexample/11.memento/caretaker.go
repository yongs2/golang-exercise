package main

type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoArray[index]
}

func (c *caretaker) getLastMemento() *memento {
	lastIndex := len(c.mementoArray) - 1
	lastMemento := c.mementoArray[lastIndex]
	c.mementoArray = c.mementoArray[:lastIndex]
	return lastMemento
}
