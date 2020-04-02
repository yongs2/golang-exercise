package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "12.golangbyexample/28.ProtocolBuffer/person"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	person := &pb.Person{Name: "XXX"}
	fmt.Printf("Person's name is %s\n", person.GetName())

	out, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Serialization error: %s", err.Error())
	}
	if err := ioutil.WriteFile("person.out", out, 0644); err != nil {
		log.Fatalf("Write File Error: %s ", err.Error())
	}
	fmt.Println("Write Success")

	in, err := ioutil.ReadFile("person.out")
	if err != nil {
		log.Fatalf("Read File Error: %s ", err.Error())
	}
	person2 := &pb.Person{}
	err2 := proto.Unmarshal(in, person2)
	if err2 != nil {
		log.Fatalf("DeSerialization error: %s", err.Error())
	}

	fmt.Println("Read Success")
	fmt.Printf("Person2's name is %s\n", person2.GetName())
}
