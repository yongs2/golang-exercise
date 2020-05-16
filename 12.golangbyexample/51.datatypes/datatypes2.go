package main

import (
	"log"
)

func main() {
	log.Printf(">> Buffered Channel\n")
	BufferedChannel()

	log.Printf(">> UnBuffered Channel\n")
	UnBufferdChannel()

	log.Printf(">> Map\n")
	Map()

	log.Printf(">> Pointer\n")
	Pointer()

	log.Printf(">> Function\n")
	Function()

	log.Printf(">> Interface\n")
	Interface()

	log.Printf(">> EmptyInterface\n")
	EmptyInterface()
}

func BufferedChannel() {
	eventsChan := make(chan string, 3)
	eventsChan <- "a"
	eventsChan <- "b"
	eventsChan <- "c"
	close(eventsChan)
	for event := range eventsChan {
		log.Printf("Event: %v\n", event)
	}
}

func UnBufferdChannel() {
	eventsChan := make(chan string)
	go sendEvents(eventsChan)
	for event := range eventsChan {
		log.Printf("Event: %v\n", event)
	}
}

func sendEvents(eventsChan chan<- string) {
	eventsChan <- "a"
	eventsChan <- "b"
	eventsChan <- "c"
	close(eventsChan)
}

func Map() {
	var employeeSalary map[string]int
	log.Printf("Map.%s: %v\n", "employeeSalary", employeeSalary)

	employeeSalary2 := make(map[string]int)
	log.Printf("Map.%s: %v\n", "employeeSalary2", employeeSalary2)

	employeeSalary3 := map[string]int{
		"John": 1000,
		"Sam":  1200,
	}
	log.Printf("Map.%s: %v\n", "employeeSalary3", employeeSalary3)

	employeeSalary3["Carl"] = 1500
	log.Printf("Map.%s: %v\n", "employeeSalary3", employeeSalary3)

	log.Printf("Map.%s: John salary is %d\n", "employeeSalary3", employeeSalary3["John"])

	delete(employeeSalary3, "Carl")
	log.Printf("Map.%s: %v\n", "employeeSalary3", employeeSalary3)
}

func Pointer() {
	var b *int
	a := 2
	b = &a

	log.Printf("%s is: %v\n", "a", a)
	log.Printf("%s is: %v\n", "b", b)
	log.Printf("%s is: %v\n", "*b", *b)
	b = new(int)
	*b = 10
	log.Printf("%s is: %v\n", "b", b)
	log.Printf("%s is: %v\n", "*b", *b)
}

func Function() {
	add := func(x, y int) int {
		return x + y
	}
	log.Printf("%s is: %d\n", "add", add(1, 2))
}

type shape interface {
	area() int
}

type square struct {
	side int
}

func (s *square) area() int {
	return s.side * s.side
}

func Interface() {
	var s shape
	s = &square{ side: 4}
	log.Printf("Interface.%s: area:%v\n", "s", s.area())
}

func EmptyInterface() {
	test("thisisstring")
	test("10")
	test(true)
}

func test(a interface{}) {
	log.Printf("%s.(%v, %T)\n", "a", a, a)
}