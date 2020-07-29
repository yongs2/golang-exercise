package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("timestamp")

	expirationTime := time.Now().Add(time.Hour)
	cookie := &http.Cookie{
		Name:     "timestamp",
		Value:    strconv.FormatInt(time.Now().Unix(), 10),
		HttpOnly: true,
		Expires:  expirationTime,
		Secure:   false,
	}
	http.SetCookie(w, cookie)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Hello, you're here for the first time")
	} else {
		timeint, _ := strconv.ParseInt(c.Value, 10, 0)
		fmt.Fprintf(w, "Hi, your last visit was at "+time.Unix(timeint, 0).Format("15:04:05"))
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
