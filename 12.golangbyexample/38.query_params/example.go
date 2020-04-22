package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	listenAddress := ":8080"

	http.Handle("/products1", http.HandlerFunc(getProducts1))
	http.Handle("/products2", http.HandlerFunc(getProducts2))
	http.Handle("/products3", http.HandlerFunc(getProducts3))
	http.Handle("/products4", http.HandlerFunc(getProducts4))

	log.Printf("Listen %v\n", listenAddress)
	http.ListenAndServe(listenAddress, nil)
}

func getProducts1(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filters, present := query["filters"]
	//log.Printf("getProducts1.fileters=[%+v],present[%+v]\n", len(filters), present)
	if !present || len(filters) == 0 {
		log.Println("filters not present")
	}
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(filters, ",")))
}

func getProducts2(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filters := query.Get("filters")
	//log.Printf("getProducts2.fileters=[%+v]\n", filters)
	w.WriteHeader(200)
	w.Write([]byte(filters))
}

func getProducts3(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filters, present := r.Form["filters"]
	//log.Printf("getProducts3.fileters=[%+v],present[%+v]\n", len(filters), present)
	if !present || len(filters) == 0 {
		log.Println("filters not present")
	}
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(filters, ",")))
}

func getProducts4(w http.ResponseWriter, r *http.Request) {
	filters := r.FormValue("filters")
	//log.Printf("getProducts4.fileters=[%+v]\n", filters)
	w.WriteHeader(200)
	w.Write([]byte(filters))
}
