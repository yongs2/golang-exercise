package main

import (
	"fmt"
)

type ele struct {
	name string
	next *ele
}

type singleList struct {
	len  int
	head *ele
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) AddFront(name string) {
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
	} else {
		ele.next = s.head
		s.head = ele
	}
	s.len++
	return
}

func (s *singleList) AddBack(name string) {
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
	} else {
		current := s.head
		if current.next != nil {
			current = current.next
		}
		current.next = ele
	}
	s.len++
	return
}

func (s *singleList) RemoveFront() error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}

func (s *singleList) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	var prev *ele
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
	return nil
}

func (s *singleList) Front() (string, error) {
	if s.head == nil {
		return "", fmt.Errorf("Single List is empty")
	}
	return s.head.name, nil
}

func (s *singleList) Size() int {
	return s.len
}

func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	current := s.head
	var result string
	for current != nil {
		if len(result) > 0 {
			result += ","
		}
		result += current.name
		current = current.next
	}
	if len(result) > 0 {
		fmt.Println("signleList:" + result)
	}
	return nil
}

func main() {
	var name string
	singleList := initList()

	name = "A"
	fmt.Printf("AddFront: %s\n", name)
	singleList.AddFront(name)
	singleList.Traverse()

	name = "B"
	fmt.Printf("AddFront: %s\n", name)
	singleList.AddFront(name)
	singleList.Traverse()

	name = "C"
	fmt.Printf("AddBack: %s\n", name)
	singleList.AddBack(name)
	singleList.Traverse()

	fmt.Printf("Size: %d\n", singleList.Size())

	err := singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveFront\n")
	err = singleList.RemoveFront()
	if err != nil {
		fmt.Printf("RemoveFront Error: %s\n", err.Error())
	}
	singleList.Traverse()

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}
	singleList.Traverse()

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}
	singleList.Traverse()

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}
	fmt.Printf("signleList: [%+v]\n", singleList)

	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Size: %d\n", singleList.Size())
}
