package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func main() {
	reqTime := time.Now()
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	rspTime := time.Now()

	fmt.Println("Response Status:", resp.Status, rspTime.Sub(reqTime))

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
