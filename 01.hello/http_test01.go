package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting 	string
	Punct 		string
	Who 		string
}

func (s String) ServeHTTP(
	w http.ResponseWriter, 
	r * http.Request ) {
	var response string
	response = fmt.Sprintf("Hello [%s]\n", s)
	fmt.Fprint(w, response)
}

func (s Struct) ServeHTTP(
	w http.ResponseWriter, 
	r * http.Request ) {
	var response string
	response = fmt.Sprintf("Hello [%s.%s.%s]\n", s.Greeting, s.Punct, s.Who)
	fmt.Fprint(w, response)
}

// Test: 
//	curl http://localhost:4000/string
//	curl http://localhost:4000/struct
func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	http.ListenAndServe("localhost:4000", nil)
}