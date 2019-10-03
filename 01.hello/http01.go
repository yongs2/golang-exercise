package main

import (
	"fmt"
	"net/http"
)

/* Package http 는 http.Handler 를 구현한 어떠 값을 사용하여 HTTP 요청(requests)을 제공합니다.
package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}
*/
type Hello struct {}

func (h Hello) ServeHTTP(
	w http.ResponseWriter, 
	r * http.Request ) {
	fmt.Fprint(w, "Hello !!!")
}

// Test: curl http://localhost:4000
func main() {
	var h Hello

	http.ListenAndServe("localhost:4000", h)
}