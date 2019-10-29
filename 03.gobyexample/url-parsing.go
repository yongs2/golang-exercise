package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgress://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(1, u.Scheme)
	fmt.Println(2, u.User)
	fmt.Println(3, u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(4, p)
	fmt.Println(5, u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(6, host)
	fmt.Println(7, port)
	fmt.Println(8, u.Path)
	fmt.Println(9, u.Fragment)
	fmt.Println(10, u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(11, m)
	fmt.Println(12, m["k"][0])
}
