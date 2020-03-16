package main

import (
	"log"
	"strconv"
)

func main() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	log.Printf("initPool=C[%v], pool.A[%v]\n", pool.capacity, len(pool.active))

	conn1, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	log.Printf("after pool.A[%v], conn1[%v]\n", len(pool.active), conn1.getID())

	conn2, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	log.Printf("after pool.A[%v], conn2[%v]\n", len(pool.active), conn2.getID())

	conn3, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	log.Printf("after pool.A[%v], conn3[%v]\n", len(pool.active), conn3.getID())

	pool.receive(conn3)
	log.Printf("after pool.A[%v], conn3[%v]\n", len(pool.active), conn3.getID())

	pool.receive(conn1)
	log.Printf("after pool.A[%v], conn1[%v]\n", len(pool.active), conn1.getID())

	pool.receive(conn2)
	log.Printf("after pool.A[%v], conn2[%v]\n", len(pool.active), conn2.getID())
}
